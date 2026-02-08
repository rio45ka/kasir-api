package models

import "time"

type Transaction struct {
	ID          string               `json:"id"`
	TotalAmount int                  `json:"total_amount"`
	CreatedAt   time.Time            `json:"created_at"`
	Details     []TransactionDetails `json:"details"`
}

type TransactionDetails struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	SubTotal      int    `json:"sub_total"`
}

type CheckoutItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}
