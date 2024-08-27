package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
