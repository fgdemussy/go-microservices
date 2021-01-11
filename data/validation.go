package data

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validate checks field validations for a given Product
func Validate(i interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(i)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-def-gjk
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches:= re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}