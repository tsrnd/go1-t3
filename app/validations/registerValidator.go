package validations

import 	"gopkg.in/go-playground/validator.v9"

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=40,alphanum"`
	Email string     `validate:"required,email"`
	Password string   `validate:"required,min=4"`
}

func (registerRequest RegisterRequest) ValidateStruct() (err error) {
	validate = validator.New()
	err = validate.Struct(registerRequest)
	return
}