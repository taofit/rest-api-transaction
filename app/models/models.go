package models

type Transaction struct {
	Transaction_id string  `json:"transaction_id"`
	Account_id     string  `json:"account_id"`
	Amount         float64 `json:"amount"`
	Created_at     string  `json:"created_at"`
}

type Account struct {
	Account_id string  `json:"account_id"`
	Balance    float64 `json:"balance"`
}
