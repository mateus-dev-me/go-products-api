package use_cases

import (
	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
)

type GetAllProductsUseCase struct {
	ProductRepo db.ProductRepositoryDB
}

func NewGetAllProductsUseCase(repo db.ProductRepositoryDB) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{ProductRepo: repo}
}

func (uc *GetAllProductsUseCase) Execute() ([]domain.Product, error) {
	return uc.ProductRepo.GetAll()
}
