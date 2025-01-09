package handler

import (
	"database/sql"
	"errors"
	"log"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// public
func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context)  {
		// TODO : get from db
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("Error when taking data from product: %v \n", err)
			c.JSON(500, gin.H{"error": "Error happened"})
			return
		}
		// TODO : do response
		c.JSON(200, products)
	}
}

func GetProduct(db *sql.DB) gin.HandlerFunc{
	return func (c *gin.Context) {
		// TODO : read id from url
		id := c.Param("id")
	
		// TODO : get from db
		product, err := model.SelectProductByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows){
				log.Printf("Error when taking data from product: %v \n", err)
				c.JSON(404, gin.H{"error": "Didn't find any product"})
				return
			}
			
			log.Printf("Error when taking data from product: %v \n", err)
			c.JSON(500, gin.H{"error": "Error happened to server"})
			return
		}
		
		// TODO : do response
		c.JSON(200, product)
	}
}



// admin
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product
		if err := c.Bind(&product); err != nil {
			log.Printf("Error when reading request body: %v \n", err)
			c.JSON(400, gin.H{"error": "Product data invalid"})
			return
		}

		product.ID = uuid.New().String()
		
		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Error when creating product: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		// TODO : do response
		c.JSON(201, product)
	}
}

func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var productReq model.Product
		if err := c.Bind(&productReq); err != nil {
			log.Printf("Error when reading request body: %v \n", err)
			c.JSON(400, gin.H{"error": "Product data invalid"})
			return
		}

		product, err := model.SelectProductByID(db, id)
		if err != nil {
			log.Printf("Error when select before update product: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		if productReq.Name != "" {
			product.Name = productReq.Name
		}
		
		if productReq.Price != 0 {
			product.Price = productReq.Price
		}

		if err := model.UpdateProduct(db, product); err != nil {
			log.Printf("Error when updating product: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		// TODO : do response
		c.JSON(201, product)
	}
}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}