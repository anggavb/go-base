package custom_validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	RegisterJsonTagNameFunc bool
}

// New creates a new instance of CustomValidator with the option to register the JSON tag name function.
//
// Parameters:
//   - registerValidator.RegisterJsonTagNameFunc [bool]: Indicates whether to register the JSON tag name function.
func New(registerValidator CustomValidator) *CustomValidator {
	customValidator := &CustomValidator{
		RegisterJsonTagNameFunc: registerValidator.RegisterJsonTagNameFunc,
	}

	customValidator.registerCustomValidator()

	return customValidator
}

func (cv *CustomValidator) registerCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if cv.RegisterJsonTagNameFunc {
			registerTagNameFunc(v)
		}
	}
}
