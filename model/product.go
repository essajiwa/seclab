package model

// product type
type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Category    string  `json:"category" db:"category"`
}
