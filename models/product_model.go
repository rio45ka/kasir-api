package models

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CreateProductRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID string `json:"category_id"`
}
