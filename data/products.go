package data

import (
	"encoding/json"
	"io"
	"time"
)

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

// FromJSON decodes json from an ioReader into a product struct
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Products is a list of Product
type Products []*Product

// ToJSON returns a JSON representation of Products
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a ref to a list of products
func GetProducts() Products {
	return productList
}

// AddProduct add a product to productList, auto-generating it's ID
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

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
