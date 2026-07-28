package custom_validator

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func registerTagNameFunc(v *validator.Validate) {
	// Register Tag Name Func
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})
	log.Println("RegisterTagNameFunc - get json/form struct tag")
}
