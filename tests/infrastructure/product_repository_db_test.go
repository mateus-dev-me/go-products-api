package infrastructure_test

import (
	"testing"

	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_Save(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)

	product := helpers.GenerateFakeProduct()

	savedID, err := repo.Save(product)
	require.NoError(t, err)
	require.NotZero(t, savedID, "Expected a non-zero ID for the saved product")

	var savedProduct domain.Product
	row := connectionDB.QueryRow(
		"SELECT id, product_name, price, category, description, quantity FROM products WHERE id = $1",
		savedID,
	)
	err = row.Scan(
		&savedProduct.ID,
		&savedProduct.Name,
		&savedProduct.Price,
		&savedProduct.Category,
		&savedProduct.Description,
		&savedProduct.Quantity,
	)
	require.NoError(t, err)

	assert.Equal(t, product.Name, savedProduct.Name)
	assert.Equal(t, product.Price, savedProduct.Price)
	assert.Equal(t, product.Category, savedProduct.Category)
	assert.Equal(t, product.Description, savedProduct.Description)
	assert.Equal(t, product.Quantity, savedProduct.Quantity)
}

func TestProductRepository_GetByID(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)

	product := helpers.GenerateFakeProduct()

	savedID, err := repo.Save(product)
	require.NoError(t, err)
	require.NotZero(t, savedID, "Expected a non-zero ID for the saved product")

	foundProduct, err := repo.GetById(savedID)
	require.NoError(t, err)

	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
	assert.Equal(t, product.Category, foundProduct.Category)
	assert.Equal(t, product.Description, foundProduct.Description)
	assert.Equal(t, product.Quantity, foundProduct.Quantity)
}

func TestProductRepository_GetAll(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)

	for i := 0; i < 10; i++ {
		product := helpers.GenerateFakeProduct()
		_, err := repo.Save(product)
		require.NoError(t, err)
	}

	products, err := repo.GetAll()
	require.NoError(t, err)

	require.Len(t, products, 10)

	for _, product := range products {
		require.NotEmpty(t, product.ID)
		require.NotEmpty(t, product.Name)
		require.NotEmpty(t, product.Price)
		require.NotEmpty(t, product.Category)
		require.NotEmpty(t, product.Description)
		require.NotEmpty(t, product.Quantity)
	}
}

func TestProductRepository_Update(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)

	productID, err := repo.Save(
		helpers.GenerateFakeProduct(),
	)

	require.NoError(t, err)
	require.NotZero(t, productID, "Expected a non-zero ID for the saved product")
	require.NoError(t, err)

	result := repo.Update(
		&domain.Product{
			Name:        "Calabresa",
			Price:       14.00,
			Category:    "Comida",
			Description: "Um pacote de calabresa com 2 unidades",
			Quantity:    30,
		},
	)

	require.NoError(t, result)
}

func TestProductRepository_Delete(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)

	productID, err := repo.Save(
		helpers.GenerateFakeProduct(),
	)

	require.NoError(t, err)
	require.NotZero(t, productID, "Expected a non-zero ID for the saved product")
	require.NoError(t, err)

	result := repo.Delete(productID)

	require.NoError(t, result)
}
