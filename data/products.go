package data

import "time"

// Product defines the structure for an API of products
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	Created     string  `json:"-"`
	Updated     string  `json:"-"`
	Deleted     string  `json:"-"`
}

// GetProducts returns a ref to a list of products
func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "tea",
		Description: "some chai tea",
		Price:       2.55,
		SKU:         "abc",
		Created:     time.Now().UTC().String(),
		Updated:     time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "latte",
		Description: "milky coffee",
		Price:       5,
		SKU:         "def",
		Created:     time.Now().UTC().String(),
		Updated:     time.Now().UTC().String(),
	},
}
