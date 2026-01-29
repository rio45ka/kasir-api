package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock of ProductRepositoryInterface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(id string) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceGetAll(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	expectedProducts := []models.Product{
		{ID: "1", Name: "P1", Price: 100, Stock: 10, CategoryID: "c1", CategoryName: "Cat1"},
		{ID: "2", Name: "P2", Price: 200, Stock: 20, CategoryID: "c1", CategoryName: "Cat1"},
	}

	mockRepo.On("GetAll").Return(expectedProducts, nil)

	products, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	mockRepo.AssertExpectations(t)
}

func TestServiceCreate(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	product := &models.Product{ID: "1", Name: "P1", CategoryID: "c1"}

	mockRepo.On("Create", product).Return(nil)

	err := service.Create(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetByID_Found(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	expectedProduct := &models.Product{ID: "1", Name: "P1", CategoryID: "c1", CategoryName: "Cat1"}

	mockRepo.On("GetByID", "1").Return(expectedProduct, nil)

	product, err := service.GetByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetByID_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("GetByID", "99").Return(nil, errors.New("Product not found"))

	product, err := service.GetByID("99")

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "Product not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestServiceUpdate(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	product := &models.Product{ID: "1", Name: "P1 Updated"}

	mockRepo.On("Update", product).Return(nil)

	err := service.Update(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceDelete(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := service.Delete("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
