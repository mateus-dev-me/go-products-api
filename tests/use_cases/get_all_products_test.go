package usecases_test

import (
	"testing"

	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/require"
)

func TestUseCase_GetAll(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	repo := db.NewProductRepositoryDB(connectionDB)
	useCase := use_cases.NewGetAllProductsUseCase(repo)

	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	err := helpers.CreateProducts(connectionDB)
	require.NoError(t, err)

	products, err := useCase.Execute()
	require.NoError(t, err)
	require.NotEmpty(t, products)
}
