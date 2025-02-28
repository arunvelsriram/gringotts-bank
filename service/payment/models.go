package payment

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	CustomerID string    `json:"customerId"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"createdAt"`
	Mode       string    `json:"mode"`
}

type Transactions []Transaction
