package repository

import (
	"seclab/model"
)

func (r *Repository) FindProductByCategory(category string) ([]model.Product, error) {
	query := "SELECT id,name,description,price,quantity,category FROM products WHERE category = '" + category + "' AND quantity > 0" // want "SQL injection vulnerability"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
