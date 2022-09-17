package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"transaction-management/app/database"
	"transaction-management/app/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

func PingPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	t1 := models.Transaction{
		Transaction_id: "0c9ab696-2a3d-4615-88cb-87133db13cf1",
		Account_id:     "0afd02d3-6c59-46e7-b7bc-992efd0b7ac2",
		Amount:         338,
		Created_at:     "2022-09-14T18:39:27.722776+00:00",
	}
	t2 := models.Transaction{
		Transaction_id: "e2da78a2-592d-41a6-be08-6f529f6e4212",
		Account_id:     "0afd02d3-6c59-46e7-b7bc-992efd0b7ac2",
		Amount:         334,
		Created_at:     "2022-09-14T18:34:13.38935+00:00",
	}
	transactions = append(transactions, t1, t2)
	c.JSON(http.StatusOK, transactions)
}

func GetSingleTransaction(c *gin.Context) {
	var transactionId = c.Param("id")
	transaction, _ := getSingleTransaction(transactionId)
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
	var transactionId = addTransaction(txI)
	if transactionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a transaction"})
		return
	}
	transaction, _ := getSingleTransaction(transactionId)
	if transaction == (models.Transaction{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func addTransaction(txI models.TransactionInput) string {
	tx, err := database.DB.Begin()
	var transactionId = generateUUID()
	if err != nil {
		return ""
	}

	_, err = tx.Exec("INSERT INTO transactions (id, account_id, amount) VALUES (?, ?, ?)",
		transactionId, txI.Account_id, txI.Amount)
	if err != nil {
		return ""
	} else {
		_, err = tx.Exec(`INSERT INTO accounts (id, balance) 
		VALUES (?, ?)
		ON CONFLICT(id) DO UPDATE SET balance = balance + ?`, txI.Account_id, txI.Amount, txI.Amount)
		if err != nil {
			return ""
		}
	}
	tx.Commit()
	return transactionId
}

func generateUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(newUUID))
}

func getSingleTransaction(transactionId string) (models.Transaction, error) {
	row := database.DB.QueryRow(`
		select id, account_id, amount, created_at from transactions where id = ?
	`, transactionId)
	var transaction models.Transaction
	if row.Err() != nil {
		return models.Transaction{}, row.Err()
	}

	err := row.Scan(&transaction.Transaction_id, &transaction.Account_id, &transaction.Amount, &transaction.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("_____", err)
		}
		return models.Transaction{}, err
	}
	return transaction, nil

}

func GetSingleAccount(c *gin.Context) {
	var account = models.Account{Account_id: "fbf4a552-2418-46c5-b308-6094ddc493a1", Balance: 13}
	c.JSON(http.StatusOK, account)
}
