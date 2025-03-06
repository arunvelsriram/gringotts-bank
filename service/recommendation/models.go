package recommendation

import "time"

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Transaction struct {
	ID         int       `json:"id"`
	CustomerID string    `json:"customerId"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"createdAt"`
	Mode       string    `json:"mode"`
}

type Transactions []Transaction

func (transactions Transactions) MonthlyTransactionAmount() float64 {
	total := 0.0
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}

func (transactions Transactions) MonthlyUpiTransactionCount() int {
	count := 0
	for _, t := range transactions {
		if t.Mode == "UPI" {
			count += 1
		}
	}
	return count
}

type Recommendation struct {
	Title       string `json:"title"`
	Product     string `json:"product"`
	Description string `json:"description"`
}

type Recommendations []Recommendation
