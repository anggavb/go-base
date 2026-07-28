package custom_validator

import (
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func registerNumericValidation(v *validator.Validate) {
	// Register Custom Validation for numeric tag
	v.RegisterValidation("numeric", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return true
		}

		value := fl.Field().String()
		if value == "" {
			return false
		}

		for _, char := range value {
			if char < '0' || char > '9' {
				return false
			}
		}
		return true
	})
	log.Println("RegisterValidation - Add custom validation for 'numeric' tag")
}
