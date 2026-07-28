package custom_validator

import (
	"log"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

func registerImageTypeValidation(v *validator.Validate) {
	// Register Custom Validator for image_type tag
	v.RegisterValidation("image_type", func(fl validator.FieldLevel) bool {
		file, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			log.Println("Invalid type for image_type validation")
			return false
		}

		// Example: Allowed types
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/bmp":  true,
			"image/heic": true,
			"image/webp": true,
		}

		return allowedTypes[file.Header.Get("Content-Type")]
	})
	log.Println("RegisterValidation - Add custom validation for 'image_type' tag")
}
