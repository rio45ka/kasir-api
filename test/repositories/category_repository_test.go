package repositories

import (
	"database/sql"
	"kasir-api/models"
	"kasir-api/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow("1", "Category 1", "Description").
		AddRow("2", "Category 2", "Description")

	// Expect query with JOIN
	mock.ExpectQuery("SELECT c.id, c.name FROM category c").WillReturnRows(rows)

	categories, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
	assert.Equal(t, "Category 1", categories[0].Name)
	assert.Equal(t, "Description", categories[0].Description)
}

func TestCreateCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	category := &models.Category{
		ID:          "1",
		Name:        "Test Category",
		Description: "Description",
	}

	mock.ExpectQuery("INSERT INTO category").
		WithArgs(category.ID, category.Name, category.Description).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	err = repo.Create(category)
	assert.NoError(t, err)
}

func TestGetCategoryByID_Found(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow("1", "Product 1", "Description")

	mock.ExpectQuery("SELECT c.id, c.name, c.description FROM category c WHERE c.id = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	category, err := repo.GetByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Product 1", category.Name)
	assert.Equal(t, "Description", category.Description)
}

func TestGetCategoryByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	mock.ExpectQuery("SELECT c.id, c.name, c.description FROM category c WHERE c.id = \\$1").
		WithArgs("99").
		WillReturnError(sql.ErrNoRows)

	category, err := repo.GetByID("99")
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, "Category not found", err.Error())
}

func TestUpdateCategory_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	category := &models.Category{
		ID:          "1",
		Name:        "Updated Category",
		Description: "Updated Description",
	}

	mock.ExpectExec("UPDATE category SET name = \\$1, description = \\$2 WHERE id = \\$3").
		WithArgs(category.Name, category.Description, category.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(category)
	assert.NoError(t, err)
}

func TestUpdateCategory_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	category := &models.Category{
		ID:          "99",
		Name:        "Updated Category",
		Description: "Updated Description",
	}

	mock.ExpectExec("UPDATE category SET name = \\$1, description = \\$2 WHERE id = \\$3").
		WithArgs(category.Name, category.Description, category.ID).
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.Update(category)
	assert.Error(t, err)
	assert.Equal(t, "Category not found", err.Error())
}

func TestDeleteCategory_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	mock.ExpectExec("DELETE FROM category WHERE id = \\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete("1")
	assert.NoError(t, err)
}

func TestDeleteCategory_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)

	mock.ExpectExec("DELETE FROM category WHERE id = \\$1").
		WithArgs("99").
		WillReturnResult(sqlmock.NewResult(1, 0))

	err = repo.Delete("99")
	assert.Error(t, err)
	assert.Equal(t, "Category not found", err.Error())
}
