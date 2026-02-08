package repositories

import (
	"database/sql"
	"kasir-api/models"
	"kasir-api/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "category_id", "category_name"}).
		AddRow("1", "Product 1", 1000, 10, "c1", "Category 1").
		AddRow("2", "Product 2", 2000, 20, "c1", "Category 1")

	// Expect query with JOIN
	mock.ExpectQuery("SELECT p.id, p.name, p.price, p.stock, c.id, c.name FROM product p JOIN category c ON p.category_id = c.id").WillReturnRows(rows)

	products, err := repo.GetAll("")
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Category 1", products[0].CategoryName)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	product := &models.Product{
		ID:         "1",
		Name:       "Test Product",
		Price:      5000,
		Stock:      10,
		CategoryID: "c1",
	}

	mock.ExpectQuery("INSERT INTO product").
		WithArgs(product.ID, product.Name, product.Price, product.Stock, product.CategoryID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	err = repo.Create(product)
	assert.NoError(t, err)
}

func TestGetByID_Found(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "category_id", "category_name"}).
		AddRow("1", "Product 1", 1000, 10, "c1", "Category 1")

	mock.ExpectQuery("SELECT p.id, p.name, p.price, p.stock, c.id, c.name FROM product p JOIN category c ON p.category_id = c.id WHERE p.id = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	product, err := repo.GetByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, "Category 1", product.CategoryName)
}

func TestGetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	mock.ExpectQuery("SELECT p.id, p.name, p.price, p.stock, c.id, c.name FROM product p JOIN category c ON p.category_id = c.id WHERE p.id = \\$1").
		WithArgs("99").
		WillReturnError(sql.ErrNoRows)

	product, err := repo.GetByID("99")
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "Product not found", err.Error())
}

func TestUpdate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	product := &models.Product{
		ID:         "1",
		Name:       "Updated Product",
		Price:      6000,
		Stock:      15,
		CategoryID: "c1",
	}

	mock.ExpectExec("UPDATE product SET name = \\$1, price = \\$2, stock = \\$3, category_id = \\$4 WHERE id = \\$5").
		WithArgs(product.Name, product.Price, product.Stock, product.CategoryID, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(product)
	assert.NoError(t, err)
}

func TestUpdate_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	product := &models.Product{
		ID:         "99",
		Name:       "Updated Product",
		Price:      6000,
		Stock:      15,
		CategoryID: "c1",
	}

	mock.ExpectExec("UPDATE product SET name = \\$1, price = \\$2, stock = \\$3, category_id = \\$4 WHERE id = \\$5").
		WithArgs(product.Name, product.Price, product.Stock, product.CategoryID, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.Update(product)
	assert.Error(t, err)
	assert.Equal(t, "Product not found", err.Error())
}

func TestDelete_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	mock.ExpectExec("DELETE FROM product WHERE id = \\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete("1")
	assert.NoError(t, err)
}

func TestDelete_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewProductRepository(db)

	mock.ExpectExec("DELETE FROM product WHERE id = \\$1").
		WithArgs("99").
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.Delete("99")
	assert.Error(t, err)
	assert.Equal(t, "Product not found", err.Error())
}
