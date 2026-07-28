package custom_validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	RegisterJsonTagNameFunc        bool
	RegisterNumericValidation      bool
	RegisterImageMaxSizeValidation bool
	RegisterImageTypeValidation    bool
}

// New creates a new instance of CustomValidator with the option to register the JSON tag name function.
//
// Parameters:
//   - registerValidator.RegisterJsonTagNameFunc [bool]: Indicates whether to register the JSON tag name function.
//   - registerValidator.RegisterNumericValidation [bool]: Indicates whether to register the numeric validation.
//   - registerValidator.RegisterImageMaxSizeValidation [bool]: Indicates whether to register the image max size validation.
//   - registerValidator.RegisterImageTypeValidation [bool]: Indicates whether to register the image type validation.
func New(registerValidator CustomValidator) *CustomValidator {
	customValidator := &CustomValidator{
		RegisterJsonTagNameFunc:        registerValidator.RegisterJsonTagNameFunc,
		RegisterNumericValidation:      registerValidator.RegisterNumericValidation,
		RegisterImageMaxSizeValidation: registerValidator.RegisterImageMaxSizeValidation,
		RegisterImageTypeValidation:    registerValidator.RegisterImageTypeValidation,
	}

	customValidator.registerCustomValidator()

	return customValidator
}

func (cv *CustomValidator) registerCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if cv.RegisterJsonTagNameFunc {
			registerTagNameFunc(v)
		}

		if cv.RegisterNumericValidation {
			registerNumericValidation(v)
		}

		if cv.RegisterImageMaxSizeValidation {
			registerImageMaxSizeValidation(v)
		}

		if cv.RegisterImageTypeValidation {
			registerImageTypeValidation(v)
		}
	}
}
