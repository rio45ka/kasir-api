package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"kasir-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductService is a mock of ProductServiceInterface
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) GetByID(id string) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandlerGetAll(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	expectedProducts := []models.Product{
		{ID: "1", Name: "P1", Price: 100, Stock: 10},
	}

	mockService.On("GetAll").Return(expectedProducts, nil)

	req, _ := http.NewRequest("GET", "/api/products", nil)
	rr := httptest.NewRecorder()

	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var products []models.Product
	err := json.NewDecoder(rr.Body).Decode(&products)
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
}

func TestHandlerCreate_Success(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	reqBody := models.CreateProductRequest{Name: "P1", Price: 100, Stock: 10}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/product", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	mockService.On("Create", mock.AnythingOfType("*models.Product")).Return(nil)

	handler.CreateProduct(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestHandlerGetByID_Success(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	expectedProduct := &models.Product{ID: "00000000-0000-0000-0000-000000000001", Name: "P1"}
	id := "00000000-0000-0000-0000-000000000001"

	mockService.On("GetByID", id).Return(expectedProduct, nil)

	req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var product models.Product
	json.NewDecoder(rr.Body).Decode(&product)
	assert.Equal(t, expectedProduct.ID, product.ID)
}

func TestHandlerGetByID_NotFound(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	id := "00000000-0000-0000-0000-000000000099"

	mockService.On("GetByID", id).Return(nil, errors.New("Product not found"))

	req, _ := http.NewRequest("GET", "/api/product/"+id, nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
