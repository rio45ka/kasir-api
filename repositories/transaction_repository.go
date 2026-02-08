package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
	"time"

	"github.com/google/uuid"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
	GetReport(startDate, endDate time.Time) (models.Report, error)
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
		CreatedAt:   time.Now(),
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetReport(startDate, endDate time.Time) (models.Report, error) {
	// Get Total Revenue & Transactions
	var totalRevenue int
	var totalTransactions int

	// Ensure endDate includes the whole day
	endDate = endDate.Add(24 * time.Hour)

	queryReport := `
		SELECT 
			COALESCE(SUM(total_amount), 0), 
			COUNT(id) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at < $2
	`
	err := repo.db.QueryRow(queryReport, startDate, endDate).Scan(&totalRevenue, &totalTransactions)
	if err != nil {
		return models.Report{}, err
	}

	if totalTransactions == 0 {
		return models.Report{}, nil
	}

	// Get Favorite Product
	var favProductName string
	var favProductQty int

	queryFavProduct := `
		SELECT 
			p.name, 
			COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN product p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	err = repo.db.QueryRow(queryFavProduct, startDate, endDate).Scan(&favProductName, &favProductQty)
	if err == sql.ErrNoRows {
		// No transactions, so no favorite product
		favProductName = ""
		favProductQty = 0
	} else if err != nil {
		return models.Report{}, err
	}

	// Construct Response
	report := models.Report{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
		FavoriteProduct: models.FavoriteProduct{
			Name:    favProductName,
			SoldQty: favProductQty,
		},
	}

	return report, nil
}
