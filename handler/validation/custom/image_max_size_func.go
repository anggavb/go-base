package custom_validator

import (
	"log"
	"mime/multipart"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func registerImageMaxSizeValidation(v *validator.Validate) {
	// Register Custom Validator for image_max_size tag
	v.RegisterValidation("image_max_size", func(fl validator.FieldLevel) bool {
		file, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			log.Println("Invalid type for image_max_size validation")
			return false
		}

		// Example: Max size 2MB
		param := fl.Param()
		log.Println(param)
		maxSize, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return false // Gagalkan validasi jika parameter tag bukan angka konkrit
		}

		if file.Size > maxSize {
			log.Println("Max size more than allowed")
			return false
		}

		return true
	})
	log.Println("RegisterValidation - Add custom validation for 'image_max_size' tag")
}
