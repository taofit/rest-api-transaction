package models

type Transaction struct {
	Transaction_id string `json:"transaction_id" binding:"uuid"`
	Account_id     string `json:"account_id" binding:"required,uuid"`
	Amount         int    `json:"amount" binding:"required,number"`
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
