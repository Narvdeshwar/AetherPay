package main

import (
	"aetherpay/shared"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionId string
	UserId        string
	Amount        float64
	Currency      string
	Status        string
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=admin password=password123 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with db", err)
	}
	db.AutoMigrate(&Transaction{})
	fmt.Println("Database connected successfully and table created successfully!")
}

func main() {
	initDB()
	app := gin.Default()
	app.POST("/api/v1/payment/process", func(c *gin.Context) {
		var req shared.PaymentRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		txnId := "1x23fdfg"

		// fmt.Printf("Processing the payment request of %s %.2f by the User: %s (%s)", req.Currency, req.Amount, req.Name, req.UserId)
		newTxn := Transaction{
			TransactionId: txnId,
			UserId:        req.UserId,
			Amount:        req.Amount,
			Currency:      req.Currency,
			Status:        "SUCCESS",
		}

		result := db.Create(&newTxn)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data doesn't strored in DB"})
			return
		}

		log.Println("Transaction saved in database")

		// Rabbit MQ
		go func() {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
			if err == nil {
				defer conn.Close()
				ch, _ := conn.Channel()
				defer ch.Close()
				// JSON message jo hum bhejenge
				msgBody := fmt.Sprintf(`{"user_id": "%s", "transaction_id": "%s", "event": "PAYMENT_SUCCESS"}`, req.UserId, txnId)
				ch.Publish(
					"",                    // Exchange
					"email_notifications", // Routing key (Queue ka naam)
					false,                 // Mandatory
					false,                 // Immediate
					amqp.Publishing{
						ContentType: "application/json",
						Body:        []byte(msgBody),
					})
				log.Println("🚀 Message sent to RabbitMQ!")
			}else {
				log.Println("⚠️ RabbitMQ connection failed:", err)
			}
		}()
		resp := shared.PaymentResponse{
			TransactionId: txnId,
			Status:        "SUCCESS",
		}
		c.JSON(http.StatusOK, resp)
	})

	fmt.Println("Server is running on the port 3002")
	app.Run(":3002")
}
