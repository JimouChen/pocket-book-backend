package models

type ParmaAddExpenses struct {
	CategoryId      int32   `json:"category_id" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	Description     string  `json:"description,omitempty"`
	Amount          float64 `json:"amount" binging:"required"`
	TransactionDate string  `json:"transaction_date" binding:"required"`
}
