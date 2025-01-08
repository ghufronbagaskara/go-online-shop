package handler

import (
	"database/sql"
	"errors"
	"log"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
)

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