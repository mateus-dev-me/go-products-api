package use_cases

import (
	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
)

type SaveProductUseCase struct {
	ProductRepo db.ProductRepositoryDB
}

func NewSaveProductUseCase(repo db.ProductRepositoryDB) *SaveProductUseCase {
	return &SaveProductUseCase{ProductRepo: repo}
}

func (uc *SaveProductUseCase) Execute(product domain.Product) (domain.Product, error) {
	productId, err := uc.ProductRepo.Save(product)
	if err != nil {
		return domain.Product{}, err
	}

	product.ID = productId

	return product, nil
}
