package usecases

import (
	"testing"

	"go-products/internal/infrastructure/db"
	"go-products/internal/use_cases"
	"go-products/tests/helpers"

	"github.com/stretchr/testify/require"
)

func TestUseCase_Delete(t *testing.T) {
	connectionDB := helpers.SetupTestDB(t)
	repo := db.NewProductRepositoryDB(connectionDB)
	useCase := use_cases.NewDeleteProductUseCase(repo)

	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	productID, err := repo.Save(
		helpers.GenerateFakeProduct(),
	)

	require.NoError(t, err)
	require.NotZero(t, productID, "Expected a non-zero ID for the saved product")
	require.NoError(t, err)

	result := useCase.Execute(productID)

	require.NoError(t, result)
}
