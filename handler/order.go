package handler

import (
	"database/sql"
	"log"
	"math/rand"
	"onlineshop/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

		// TODO: make passcode
		passcode := generatePasccode(5)


		// TODO: hash passcode
		hashcode, err := bcrypt.GenerateFromPassword([]byte(passcode), 10)
		if err != nil {
			log.Printf("Error occur when making hash in checkout order : %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		hascodeString := string(hashcode)


		// TODO: make order & detail
		order := model.Order {
			ID: uuid.New().String(),
			Email: checkoutOrder.Email,
			Address: checkoutOrder.Address,
			Passcode: &hascodeString,
			GrandTotal: 0,
		}

		details := []model.OrderDetail{}

		for _,p := range products {
			total := p.Price * int64(orderQty[p.ID])

			detail := model.OrderDetail{
				ID: uuid.New().String(),
				OrderID: order.ID,
				ProductID: p.ID,
				Quantity: orderQty[p.ID],
				Price: p.Price,
				Total: total,
			}

			details = append(details, detail)

			order.GrandTotal += total
		}

		model.CreateOrder(db, order, details)

		orderWithDetail := model.OrderWithDetail{
			Order: order,
			Details: details,
		}

		orderWithDetail.Order.Passcode = &passcode
		
		c.JSON(200, orderWithDetail)


	}
}

func generatePasccode(length int) string {
	charset := "ABCDEFGHIJKLMNOPQESTUVWXYZ1234567890"

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[randomGenerator.Intn(len(charset))]
	}

	return string(code)
}

func ConfirmOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: retrieve id from param
		id := c.Param("id")

		//TODO: parse req body
		var confirmReq model.Confirm
		if err := c.BindJSON(&confirmReq); err != nil {
			log.Printf("Error when parsing request body: %v \n", err)
			c.JSON(400, gin.H{"error": "Order data invalid"})
			return
		}

		//TODO: fetch order data from db
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Error when fetching order data: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		//TODO: verify passcode
		if order.Passcode == nil {
			log.Println("Passcode unvalid at confirm order")
			c.JSON(400, gin.H{"error": "Order data invalid"})
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(confirmReq.Passcode)); err != nil {
			log.Printf("Error when verify passcode: %v \n", err)
			c.JSON(401, gin.H{"error": "Unallowed to access order"})
			return
		}

		//TODO: check the order isnt pay yet
		if order.PaidAt != nil {
			log.Println("Order already paid")
			c.JSON(400, gin.H{"error": "Order already paid"})
			return
		}

		//TODO: verify payment amount
		if order.GrandTotal != confirmReq.Amount {
			log.Printf("Amount isnt match: %v \n", err)
			c.JSON(400, gin.H{"error": "Payment amount invalid"})
			return
		}

		//TODO: update and confirm the order
		current := time.Now()
		if err = model.UpdateOrderById(db, id, confirmReq, current); err != nil {
			log.Printf("Error when updating order data: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		order.Passcode = nil
		
		order.PaidAt = &current
		order.PaidBank = &confirmReq.Bank
		order.PaidAccountNumber = &confirmReq.AccountNumber
		
		c.JSON(200, order)
	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: retrieve id from param
		id := c.Param("id")

		//TODO: retrieve passcode from query parameter
		passcode := c.Query("passcode")

		//TODO: fetch order data from db
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Error when fetching order data: %v \n", err)
			c.JSON(500, gin.H{"error": "Server error"})
			return
		}

		//TODO: verify passcode
		if order.Passcode == nil {
			log.Println("Passcode unvalid at confirm order")
			c.JSON(400, gin.H{"error": "Order data invalid"})
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(passcode)); err != nil {
			log.Printf("Error when verify passcode: %v \n", err)
			c.JSON(401, gin.H{"error": "Unallowed to access order"})
			return
		}
		
		order.Passcode = nil
		c.JSON(200, order)
	}
}