package helpers

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"

	"github.com/bxcodec/faker/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

const (
	testDBFile = "test.db"
)

func SetupTestDB(t *testing.T) *sql.DB {
	dbFile := filepath.Join(os.TempDir(), testDBFile)
	db, err := sql.Open("sqlite3", dbFile)
	require.NoError(t, err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_name TEXT NOT NULL,
			price REAL NOT NULL,
			category TEXT NOT NULL,
			description TEXT,
			quantity INTEGER NOT NULL
		)
	`)
	require.NoError(t, err)

	return db
}

func TeardownTestDB(t *testing.T) {
	dbFile := filepath.Join(os.TempDir(), testDBFile)
	err := os.Remove(dbFile)
	require.NoError(t, err)
}

func GenerateFakeProduct() domain.Product {
	name := faker.Word()
	category := faker.Word()
	description := faker.Paragraph()

	priceSlice, err := faker.RandomInt(1, 100)
	if err != nil {
		log.Fatalf("Erro ao gerar preço fictício: %v", err)
	}
	quantitySlice, err := faker.RandomInt(1, 100)
	if err != nil {
		log.Fatalf("Erro ao gerar quantidade fictícia: %v", err)
	}

	price := float64(priceSlice[0]) // Converte para float64
	quantity := quantitySlice[0]

	return domain.Product{
		Name:        name,
		Price:       price,
		Category:    category,
		Description: description,
		Quantity:    quantity,
	}
}

func CreateProduct(connection *sql.DB) error {
	repo := db.NewProductRepositoryDB(connection)

	product := GenerateFakeProduct()
	_, err := repo.Save(product)
	if err != nil {
		return err
	}

	return nil
}

func CreateProducts(connection *sql.DB) error {
	repo := db.NewProductRepositoryDB(connection)

	for i := 0; i < 10; i++ {
		product := GenerateFakeProduct()
		_, err := repo.Save(product)
		if err != nil {
			return err
		}
	}

	return nil
}
