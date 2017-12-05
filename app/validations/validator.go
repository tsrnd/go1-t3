package validations

import 	"gopkg.in/go-playground/validator.v9"

var validate *validator.Validate

type Validator interface {
	ValidateStruct()
}
