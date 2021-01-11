package data

import (
	"fmt"
)

// ErrProductNotFound denotes when a product can not be found in datastore
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API of products
// swagger:model
// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products is a slice of Product
type Products []*Product

// GetProducts returns a ref to a list of products
func GetProducts() Products {
	return productList
}

// AddProduct add a product to productList, auto-generating it's ID
func AddProduct(p Product) {
	p.ID = getNextID()
	productList = append(productList, &p)
}

// UpdateProduct replaces product with given data
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	productList[i] = &p

	return nil
}

// GetProductByID returns a single product which matches the id from the
// datastore.
// Returns ProductNotFound error
func GetProductByID(id int) (*Product, error)  {
	i := findIndexByProductID(id);
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// getNextID returns the next ID in sequence
func getNextID() int {
	lp := productList[len(productList) -1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "tea",
		Description: "some chai tea",
		Price:       2.55,
		SKU:         "abc",
	},
	{
		ID:          2,
		Name:        "latte",
		Description: "milky coffee",
		Price:       5,
		SKU:         "def",
	},
}
