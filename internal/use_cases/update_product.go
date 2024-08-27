package use_cases

import (
	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
)

type UpdateProductUseCase struct {
	ProductRepo db.ProductRepositoryDB
}

func NewUpdateProductUseCase(repo db.ProductRepositoryDB) *UpdateProductUseCase {
	return &UpdateProductUseCase{ProductRepo: repo}
}

func (uc *UpdateProductUseCase) Execute(product *domain.Product) error {
	return uc.ProductRepo.Update(product)
}
