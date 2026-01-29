package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository is a mock of CategoryRepositoryInterface
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAll() ([]models.Category, error) {
	args := m.Called()
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Create(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(id string) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceCategoryGetAll(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	expectedCategories := []models.Category{
		{ID: "1", Name: "Cat1", Description: "Description1"},
		{ID: "2", Name: "Cat2", Description: "Description2"},
	}

	mockRepo.On("GetAll").Return(expectedCategories, nil)

	categories, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
	mockRepo.AssertExpectations(t)
}

func TestServiceCategoryCreate(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	category := &models.Category{ID: "1", Name: "Cat1", Description: "Description1"}

	mockRepo.On("Create", category).Return(nil)

	err := service.Create(category)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceCategoryGetByID_Found(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	expectedCategory := &models.Category{ID: "1", Name: "Cat1", Description: "Description1"}

	mockRepo.On("GetByID", "1").Return(expectedCategory, nil)

	category, err := service.GetByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
	mockRepo.AssertExpectations(t)
}

func TestServiceCategoryGetByID_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	mockRepo.On("GetByID", "99").Return(nil, errors.New("Category not found"))

	category, err := service.GetByID("99")

	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, "Category not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestServiceCategoryUpdate(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	category := &models.Category{ID: "1", Name: "Cat1 Updated", Description: "Description1 Updated"}

	mockRepo.On("Update", category).Return(nil)

	err := service.Update(category)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestServiceCategoryDelete(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := services.NewCategoryService(mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := service.Delete("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
