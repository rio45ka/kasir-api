package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type CategoryHandler struct {
	service services.CategoryServiceInterface
}

func NewCategoryHandler(service services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategoryByID - GET/PUT/DELETE /api/category/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

// GET All Categories
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := h.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error(http.StatusInternalServerError, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Categories retrieved successfully", categories))
}

// Create Category
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	category := models.Category{
		ID:   uuid.NewString(),
		Name: req.Name,
	}

	err = h.service.Create(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Success(http.StatusCreated, "Category created successfully", category))
}

// GET Category By ID
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrCategoryNotFound, err.Error()))
		return
	}

	category, err := h.service.GetByID(id.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Error(http.StatusNotFound, models.ErrCategoryNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Category retrieved successfully", category))
}

// Update Category
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrCategoryNotFound, err.Error()))
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	category.ID = id.String()
	err = h.service.Update(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Category updated successfully", category))
}

// Delete Category
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrCategoryNotFound, err.Error()))
		return
	}

	err = h.service.Delete(id.String())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Success(http.StatusOK, "Category deleted successfully", nil))
}
