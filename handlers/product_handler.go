package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type ProductHandler struct {
	service services.ProductServiceInterface
}

func NewProductHandler(service services.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProductByID - GET/PUT/DELETE /api/produk/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET All Products
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error(http.StatusInternalServerError, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Products retrieved successfully", products))
}

// Create Product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	product := models.Product{
		ID:         uuid.NewString(),
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	err = h.service.Create(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Success(http.StatusCreated, "Product created successfully", product))
}

// GET Product By ID
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrProductNotFound, err.Error()))
		return
	}

	product, err := h.service.GetByID(id.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Error(http.StatusNotFound, models.ErrProductNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Product retrieved successfully", product))
}

// Update Product
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrProductNotFound, err.Error()))
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrProductNotFound, err.Error()))
		return
	}

	product.ID = id.String()
	err = h.service.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Product updated successfully", product))
}

// Delete Product
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrProductNotFound, err.Error()))
		return
	}

	err = h.service.Delete(id.String())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrProductNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Product deleted successfully", nil))
}
