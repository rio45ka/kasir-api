package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"

	"github.com/google/uuid"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
}

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {

	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetails, 0)

	// LOOP VALIDASI & CALCULATE TOTAL
	for _, item := range items {

		var productPrice, stock int
		var productName string

		err := tx.QueryRow(
			"SELECT name, price, stock FROM product WHERE id = $1",
			item.ProductID,
		).Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %s not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("stock not enough for product %s", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// update stock
		_, err = tx.Exec(
			"UPDATE product SET stock = stock - $1 WHERE id = $2",
			item.Quantity,
			item.ProductID,
		)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetails{
			ID:          uuid.NewString(),
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			SubTotal:    subtotal,
		})
	}

	// INSERT TRANSACTION
	var transactionID string
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id",
		totalAmount,
	).Scan(&transactionID)

	if err != nil {
		return nil, err
	}

	// 🔥 INSERT DETAILS
	for i := range details {
		details[i].TransactionID = transactionID

		_, err = tx.Exec(
			"INSERT INTO transaction_details (id, transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4, $5)",
			details[i].ID,
			transactionID,
			details[i].ProductID,
			details[i].Quantity,
			details[i].SubTotal,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
