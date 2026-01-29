package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryServiceInterface interface {
	GetAll() ([]models.Category, error)
	Create(data *models.Category) error
	GetByID(id string) (*models.Category, error)
	Update(product *models.Category) error
	Delete(id string) error
}

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *models.Category) error {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(product *models.Category) error {
	return s.repo.Update(product)
}

func (s *CategoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
