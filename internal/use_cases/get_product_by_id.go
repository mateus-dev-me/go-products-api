package use_cases

import (
	"go-products/internal/domain"
	"go-products/internal/infrastructure/db"
)

type GetProductByIDUseCase struct {
	ProductRepo db.ProductRepositoryDB
}

func NewGetByIDProductUseCase(repo db.ProductRepositoryDB) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{ProductRepo: repo}
}

func (uc *GetProductByIDUseCase) Execute(id int) (*domain.Product, error) {
	product, err := uc.ProductRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
