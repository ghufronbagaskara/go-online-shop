package handler

import (
	"database/sql"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func (c *gin.Context)  {
		// TODO : get from db
		products, err := model.SelectProduct(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error happened"})
			return
		}
		// TODO : do response
		c.JSON(200, products)
	}
}

func GetProduct(c *gin.Context){
	// TODO : read id from url

	// TODO : get from db
	
	// TODO : do response
}