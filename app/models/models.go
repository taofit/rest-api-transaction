package models

import (
	"database/sql"
	"log"
	"os/exec"
	"strings"
	"transaction-management/app/database"
)

type Transaction struct {
	Transaction_id string `json:"transaction_id" binding:"required,uuid"`
	Account_id     string `json:"account_id" binding:"required,uuid"`
	Amount         int    `json:"amount" binding:"number"`
	Created_at     string `json:"created_at"`
}

type Account struct {
	Account_id string `json:"account_id" binding:"required,uuid"`
	Balance    int    `json:"balance" binding:"number"`
}

type TransactionInput struct {
	Account_id string `json:"account_id" binding:"required,uuid"`
	Amount     int    `json:"amount" binding:"required,number"`
}

func GetTransactions() ([]Transaction, error) {
	rows, err := database.DB.Query("SELECT id, account_id, amount, created_at FROM transactions")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var transactions []Transaction
	for rows.Next() {
		transaction := Transaction{}
		err = rows.Scan(&transaction.Transaction_id, &transaction.Account_id, &transaction.Amount, &transaction.Created_at)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func AddTransaction(txI TransactionInput) string {
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
		ON CONFLICT(id) DO UPDATE SET balance = balance + ?`,
			txI.Account_id, txI.Amount, txI.Amount)
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

func GetSingleTransaction(transactionId string) (Transaction, error) {
	row := database.DB.QueryRow("SELECT id, account_id, amount, created_at FROM transactions WHERE id = ?", transactionId)
	if row.Err() != nil {
		return Transaction{}, row.Err()
	}
	var transaction Transaction
	err := row.Scan(&transaction.Transaction_id, &transaction.Account_id, &transaction.Amount, &transaction.Created_at)
	if err != nil && err == sql.ErrNoRows {
		return Transaction{}, nil
	}
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func GetSingleAccount(accountId string) (Account, error) {
	row := database.DB.QueryRow("SELECT id, balance FROM accounts WHERE id = ?", accountId)
	if row.Err() != nil {
		return Account{}, row.Err()
	}
	var account Account
	err := row.Scan(&account.Account_id, &account.Balance)
	if err != nil && err == sql.ErrNoRows {
		return Account{}, nil
	}
	if err != nil {
		return Account{}, err
	}
	return account, nil
}
