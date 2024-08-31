package usecases_test

import (
	"testing"

	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/require"
)

func TestUsaseCase_Save(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	repo := db.NewProductRepositoryDB(connectionDB)
	useCase := use_cases.NewSaveProductUseCase(repo)

	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	productID, err := useCase.Execute(
		helpers.GenerateFakeProduct(),
	)

	require.NoError(t, err)
	require.NotZero(t, productID)
}
