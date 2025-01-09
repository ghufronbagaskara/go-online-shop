package handler

import (
	"database/sql"
	"log"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
)

func CheckoutOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: take order data from request 
		var checkoutOrder model.Checkout
		if err := c.BindJSON(&checkoutOrder); err != nil {
			log.Printf("Error when reading request body: %v \n", err)
			c.JSON(400, gin.H{"error": "Product data invalid"})
			return
		}

		ids := []string{}
		orderQty := make(map[string]int32)
		for _, o := range checkoutOrder.Products{
			ids = append(ids, o.ID)
			orderQty[o.ID] = o.Quantity
		}
		// TODO: take produdct from db
		products, err := model.SelectProductIn(db, ids)
		if err != nil {
			log.Printf("Error occur when retrieving product data : %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		c.JSON(200, products)

		// TODO: make passcode


		// TODO: hash passcode


		// TODO: make order & detail


	}
}

func ConfirmOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}