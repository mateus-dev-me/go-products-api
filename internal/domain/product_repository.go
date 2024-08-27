package domain

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	Save(product Product) (int, error)
	Update(product *Product) error
	Delete(id int) error
}
