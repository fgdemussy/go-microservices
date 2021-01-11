package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "hey",
		Price: 1.00,
		SKU: "abc-def-sdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}