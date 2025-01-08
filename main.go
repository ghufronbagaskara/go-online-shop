package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

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

	// 3. running server
	server := &http.Server{
		Addr: ".8080",
		Handler: nil,
	}
	if err = server.ListenAndServe(); err != nil{
		fmt.Printf("Error running server: %v \n", err)
		os.Exit(1)
	}


}