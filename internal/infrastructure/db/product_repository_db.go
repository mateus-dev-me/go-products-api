package db

import (
	"database/sql"
	"fmt"

	"go-products/internal/domain"
)

type ProductRepositoryDB struct {
	connection *sql.DB
}

func NewProductRepositoryDB(connection *sql.DB) ProductRepositoryDB {
	return ProductRepositoryDB{
		connection: connection,
	}
}

func (repo *ProductRepositoryDB) GetAll() ([]domain.Product, error) {
	query := `
		SELECT id, product_name, price, description, category, quantity FROM products
	`
	rows, err := repo.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []domain.Product{}, nil
	}
	defer rows.Close()

	var productList []domain.Product
	var productObj domain.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
			&productObj.Description,
			&productObj.Category,
			&productObj.Quantity)
		if err != nil {
			fmt.Println(err)
			return []domain.Product{}, nil
		}

		productList = append(productList, productObj)
	}

	return productList, nil
}

func (repo *ProductRepositoryDB) GetById(productId int) (*domain.Product, error) {
	row, err := repo.connection.Prepare(`
		SELECT id, product_name, price, description, category, quantity  
		FROM products 
		WHERE id=$1`,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer row.Close()

	var product domain.Product
	err = row.QueryRow(productId).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Category,
		&product.Quantity,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepositoryDB) Save(product domain.Product) (int, error) {
	var id int
	query, err := repo.connection.Prepare(`
		INSERT INTO products
		(product_name, price, description, category, quantity ) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
	)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(product.Name, product.Price, product.Description, product.Category, product.Quantity).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (repo *ProductRepositoryDB) Update(product *domain.Product) error {
	query := `UPDATE products SET product_name = $1, price = $2 WHERE id = $3`
	_, err := repo.connection.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (repo *ProductRepositoryDB) Delete(productId int) error {
	query := `DELETE FROM products WHERE id=$1 `
	_, err := repo.connection.Exec(query, productId)
	return err
}
