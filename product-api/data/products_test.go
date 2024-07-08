package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Latte",
		Price: 1.00,
		SKU:   "abc-def-ghi",
	}

	err := p.Validate()

	if err != nil {
		t.Fatalf("Product struct is not valid. Got: %s", err)
	}
}
