package usecases_test

import (
	"testing"

	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/require"
)

func TestUseCase_Update(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	repo := db.NewProductRepositoryDB(connectionDB)
	useCase := use_cases.NewUpdateProductUseCase(repo)

	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	productID, err := repo.Save(
		helpers.GenerateFakeProduct(),
	)

	require.NoError(t, err)
	require.NotZero(t, productID, "Expected a non-zero ID for the saved product")
	require.NoError(t, err)

	result := useCase.Execute(
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
