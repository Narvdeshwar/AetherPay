package main

import (
	"aetherpay/shared"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.POST("/api/v1/billing/subscribe", func(c *gin.Context) {
		paymentData := shared.PaymentRequest{
			Name:     "Ashrith",
			Amount:   2001.4,
			Currency: "INR",
			UserId:   "123",
		}
		// convert data in json
		reqData, err := json.Marshal(paymentData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert into the json"})
			return
		}
		paymentURL := "http://localhost:3002/api/v1/payment/process"
		resp, err := http.Post(paymentURL, "application/json", bytes.NewBuffer(reqData))
		log.Println("Current resposnc from payments", resp, err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load data"})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		var paymentResponse shared.PaymentResponse
		json.Unmarshal(body, &paymentResponse)
		log.Println(paymentResponse)
		if paymentResponse.Status == "SUCCESS" {
			c.JSON(http.StatusOK, gin.H{
				"message":        "subscription sucessfull",
				"transaction_id": paymentResponse.TransactionId,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":        "payment failed",
				"transaction_id": paymentResponse.TransactionId,
			})
		}

	})
	fmt.Println("Billing is runnig on the port 3001")
	app.Run(":3001")
}
