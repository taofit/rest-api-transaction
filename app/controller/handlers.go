package controller

import (
	"fmt"
	"net/http"
	"transaction-management/app/models"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetTransactions(c *gin.Context) {
	transactions, err := models.GetTransactions()
	if checkErr(c, err) {
		return
	}

	if err != nil || transactions == nil {
		c.JSON(http.StatusOK, gin.H{"error": "Transactions not found"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetSingleTransaction(c *gin.Context) {
	var transactionId = c.Param("id")
	transaction, err := models.GetSingleTransaction(transactionId)
	if checkErr(c, err) {
		return
	}

	if transaction == (models.Transaction{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func AddTransaction(c *gin.Context) {
	var txI models.TransactionInput
	if !isRequestHeaderValid(c) {
		return
	}
	if !isRequestBodyValid(c, &txI) {
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

	c.JSON(http.StatusCreated, transaction)
}

func GetSingleAccount(c *gin.Context) {
	var accountId = c.Param("id")
	account, err := models.GetSingleAccount(accountId)
	if checkErr(c, err) {
		return
	}
	if account == (models.Account{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func isRequestBodyValid(c *gin.Context, txI *models.TransactionInput) bool {
	if err := c.ShouldBindJSON(txI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mandatory body parameters missing or have incorrect type."})
		return false
	}
	return true
}

func isRequestHeaderValid(c *gin.Context) bool {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Specified content type not allowed."})
		return false
	}
	return true
}

func checkErr(c *gin.Context, err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return true
	}
	return false
}
