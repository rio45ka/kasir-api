package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"kasir-api/handlers"
	"kasir-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryService is a mock of CategoryServiceInterface
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetAll() ([]models.Category, error) {
	args := m.Called()
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) Create(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryService) GetByID(id string) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) Update(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandlerCategoryGetAll(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := handlers.NewCategoryHandler(mockService)

	expectedCategories := []models.Category{
		{ID: "1", Name: "C1"},
	}

	mockService.On("GetAll").Return(expectedCategories, nil)

	req, _ := http.NewRequest("GET", "/api/categories", nil)
	rr := httptest.NewRecorder()

	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var categories []models.Category
	err := json.NewDecoder(rr.Body).Decode(&categories)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
}

func TestHandlerCategoryCreate_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := handlers.NewCategoryHandler(mockService)

	reqBody := models.CreateCategoryRequest{Name: "C1"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/category", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	mockService.On("Create", mock.AnythingOfType("*models.Category")).Return(nil)

	handler.CreateCategory(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestHandlerCategoryGetByID_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := handlers.NewCategoryHandler(mockService)

	expectedCategory := &models.Category{ID: "00000000-0000-0000-0000-000000000001", Name: "C1"}
	id := "00000000-0000-0000-0000-000000000001"

	mockService.On("GetByID", id).Return(expectedCategory, nil)

	req, _ := http.NewRequest("GET", "/api/category/"+id, nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var category models.Category
	json.NewDecoder(rr.Body).Decode(&category)
	assert.Equal(t, expectedCategory.ID, category.ID)
}

func TestHandlerCategoryGetByID_NotFound(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := handlers.NewCategoryHandler(mockService)

	id := "00000000-0000-0000-0000-000000000099"

	mockService.On("GetByID", id).Return(nil, errors.New("Category not found"))

	req, _ := http.NewRequest("GET", "/api/category/"+id, nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
