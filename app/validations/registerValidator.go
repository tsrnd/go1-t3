package validations

import 	"gopkg.in/go-playground/validator.v9"

type RegisterRequest struct {
	username string `validate:"required,min=3,max=40,regexp=^[a-zA-Z]*$"`
	email string     `validate:"required,email"`
	password string   `validate:"required, min=6"`
}

func (registerRequest RegisterRequest) ValidateStruct() (err error) {
	validate = validator.New()
	err = validate.Struct(registerRequest)
	return
}