package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionServiceInterface interface {
	Checkout(items []models.CheckoutItem) (*models.Transaction, error)
}

type TransactionService struct {
	repository repositories.TransactionRepositoryInterface
}

func NewTransactionService(repository repositories.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{repository: repository}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repository.CreateTransaction(items)
}
