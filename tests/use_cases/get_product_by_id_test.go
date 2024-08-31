package usecases_test

import (
	"testing"

	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/require"
)

func TestUseCase_GetById(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	repo := db.NewProductRepositoryDB(connectionDB)
	useCase := use_cases.NewGetByIDProductUseCase(repo)

	product := helpers.GenerateFakeProduct()
	productID, err := repo.Save(product)

	require.NoError(t, err)
	require.NotZero(t, productID)

	result, err := useCase.Execute(productID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}
