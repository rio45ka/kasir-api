package models

type Report struct {
	TotalRevenue      int             `json:"total_revenue"`
	TotalTransactions int             `json:"total_transactions"`
	FavoriteProduct   FavoriteProduct `json:"favorite_product"`
}

type FavoriteProduct struct {
	Name    string `json:"name"`
	SoldQty int    `json:"sold_qty"`
}

type ReportResponse struct {
	Report Report `json:"report"`
}
