package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductServiceInterface interface {
	GetAll(name string) ([]models.Product, error)
	Create(data *models.Product) error
	GetByID(id string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
}

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) Create(data *models.Product) error {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
