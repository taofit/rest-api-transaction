package controller

import (
	"log"
	"net/http"
	"transaction-management/app/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

func PingPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetTransactions(c *gin.Context) {
	transactions, err := models.GetTransactions()
	checkErr(err)

	if err != nil || transactions == nil {
		c.JSON(http.StatusOK, gin.H{"error": "Transactions not found"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetSingleTransaction(c *gin.Context) {
	var transactionId = c.Param("id")
	transaction, err := models.GetSingleTransaction(transactionId)
	checkErr(err)

	if transaction == (models.Transaction{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func AddTransaction(c *gin.Context) {
	var txI models.TransactionInput
	if err := c.ShouldBindJSON(&txI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var transactionId = models.AddTransaction(txI)
	if transactionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a transaction"})
		return
	}
	transaction, _ := models.GetSingleTransaction(transactionId)
	if transaction == (models.Transaction{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch created transaction, it cannot be found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func GetSingleAccount(c *gin.Context) {
	var accountId = c.Param("id")
	account, err := models.GetSingleAccount(accountId)
	checkErr(err)

	if account == (models.Account{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
