package use_cases

import "go-products/internal/infrastructure/db"

type DeleteProductUseCase struct {
	ProductRepo db.ProductRepositoryDB
}

func NewDeleteProductUseCase(repo db.ProductRepositoryDB) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepo: repo}
}

func (uc *DeleteProductUseCase) Execute(id int) error {
	return uc.ProductRepo.Delete(id)
}
