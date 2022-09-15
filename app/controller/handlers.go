package controller

import (
	"net/http"
	"transaction-management/app/models"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	t1 := models.Transaction{
		Transaction_id: "0c9ab696-2a3d-4615-88cb-87133db13cf1",
		Account_id:     "0afd02d3-6c59-46e7-b7bc-992efd0b7ac2",
		Amount:         33.8,
		Created_at:     "2022-09-14T18:39:27.722776+00:00",
	}
	t2 := models.Transaction{
		Transaction_id: "e2da78a2-592d-41a6-be08-6f529f6e4212",
		Account_id:     "0afd02d3-6c59-46e7-b7bc-992efd0b7ac2",
		Amount:         337.4,
		Created_at:     "2022-09-14T18:34:13.38935+00:00",
	}
	transactions = append(transactions, t1, t2)
	c.JSON(http.StatusOK, transactions)
}

func GetSingleTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"account_id":     "0afd02d3-6c59-46e7-b7bc-992c5e0b7ac2",
		"amount":         77.7,
		"transaction_id": "b9c4da39-d795-4094-a880-56155f253d01",
		"created_at":     "2022-09-14T18:33:30.733082+00:00",
	})
}

func AddTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transaction_id": "1a3c9bde-de3b-4f9d-ad73-684ce49990c3",
		"account_id":     "0afd02d3-6c59-46e7-b7bc-992efd0b7ac2",
		"amount":         -11.1,
	})
}

func GetSingleAccount(c *gin.Context) {
	var account = models.Account{Account_id: "fbf4a552-2418-46c5-b308-6094ddc493a1", Balance: 13.4}
	c.JSON(http.StatusOK, account)
}
