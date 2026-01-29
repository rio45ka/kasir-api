package handlers

import (
	"encoding/json"
	"net/http"
)

type MetaHandler struct{}

func NewMetaHandler() *MetaHandler {
	return &MetaHandler{}
}

func (h *MetaHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"status":  "OK",
		"message": "Kasir API is running",
	}

	json.NewEncoder(w).Encode(response)
}

func (h *MetaHandler) ListAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"status":  "OK",
		"message": "Kasir API is running",
		"endpoints": []map[string]string{
			{
				"name":        "Health Check",
				"method":      "GET",
				"url":         "/health",
				"description": "Check if the API is running",
			},
			{
				"name":        "Get All Products",
				"method":      "GET",
				"url":         "/api/products",
				"description": "Retrieve all products",
			},
			{
				"name":        "Get Product By ID",
				"method":      "GET",
				"url":         "/api/product/{id}",
				"description": "Retrieve product by UUID",
			},
			{
				"name":        "Create Product",
				"method":      "POST",
				"url":         "/api/product",
				"description": "Create a new product",
			},
			{
				"name":        "Update Product",
				"method":      "PUT",
				"url":         "/api/product/{id}",
				"description": "Update product by UUID",
			},
			{
				"name":        "Delete Product",
				"method":      "DELETE",
				"url":         "/api/product/{id}",
				"description": "Delete product by UUID",
			},
			{
				"name":        "Get All Categories",
				"method":      "GET",
				"url":         "/api/categories",
				"description": "Retrieve all categories",
			},
			{
				"name":        "Get Category By ID",
				"method":      "GET",
				"url":         "/api/category/{id}",
				"description": "Retrieve category by UUID",
			},
			{
				"name":        "Create Category",
				"method":      "POST",
				"url":         "/api/category",
				"description": "Create a new category",
			},
			{
				"name":        "Update Category",
				"method":      "PUT",
				"url":         "/api/category/{id}",
				"description": "Update category by UUID",
			},
			{
				"name":        "Delete Category",
				"method":      "DELETE",
				"url":         "/api/category/{id}",
				"description": "Delete category by UUID",
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}
