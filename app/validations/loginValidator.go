package validations

import 	"gopkg.in/go-playground/validator.v9"

type LoginRequest struct {
	Email string     `validate:"required,email"`
	Password string   `validate:"required"`
}

func (LoginRequest LoginRequest) ValidateStruct() (err error) {
	validate = validator.New()
	err = validate.Struct(LoginRequest)
	return
}