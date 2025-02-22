package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"onlineshop/handler"
	"onlineshop/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
) 

	

func main() {
	// 1. connect to db
	db, err := sql.Open("pgx", os.Getenv("DB_URI")) 
	if err != nil {
		fmt.Printf("Error connect to db: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()


	if err = db.Ping(); err != nil {
		fmt.Printf("Error to verif db connection: %v\n", err)
		os.Exit(1)
	}

	// 2. table migration
	if _, err = migrate(db); err != nil {
		fmt.Printf("Error migrating data: %v \n", err)
		os.Exit(1)
	}

	// 3. creating handler for routing
	r := gin.Default()

	r.GET("/api/v1/products", handler.ListProducts(db))
	r.GET("/api/v1/products/:id", handler.GetProduct(db))
	r.POST("/api/v1/checkout", handler.CheckoutOrder(db))

	r.POST("/api/v1/orders/:id/confirm", handler.ConfirmOrder(db))
	r.GET("/api/v1/orders/:id", handler.GetOrder(db))

	r.POST("/admin/products", middleware.AdminOnly(), handler.CreateProduct(db))
	r.PUT("/admin/products/:id", middleware.AdminOnly(), handler.UpdateProduct(db))
	r.DELETE("/admin/products/:id", middleware.AdminOnly(), handler.DeleteProduct(db))

	// 4. running server
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	if err = server.ListenAndServe(); err != nil{
		fmt.Printf("Error running server: %v \n", err)
		os.Exit(1)
	}
}