package models

type ParmaAddExpenses struct {
	CategoryId      int32   `json:"category_id" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	Description     string  `json:"description,omitempty"`
	Amount          float64 `json:"amount" binging:"required"`
	TransactionDate string  `json:"transaction_date" binding:"required"`
	Type            int8    `json:"type" db:"type" binding:"required"`
}

type ParamSearchExpenses struct {
	// 可选参数来控制
	Title                string `json:"title,omitempty" db:"title"`
	Type                 int8   `json:"type" db:"type" binding:"required"`
	TransactionBeginDate string `json:"transaction_begin_date,omitempty"`
	TransactionEndDate   string `json:"transaction_end_date,omitempty"`
	Limit                int    `json:"limit" binding:"required"`
	Offset               int    `json:"offset" binding:"required"`
}

type ResponseSearchExpenses struct {
	Cate            string  `json:"cate" db:"cate"`
	Title           string  `json:"title" db:"title"`
	Description     string  `json:"description" db:"description"`
	Amount          float64 `json:"amount" db:"amount"`
	TransactionDate string  `json:"date" db:"date"`
}

type ResponseSearchList struct {
	// 定义你的结构体字段
}
