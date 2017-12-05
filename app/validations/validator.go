package validations

import 	"gopkg.in/go-playground/validator.v9"
import 	"strings"

var validate *validator.Validate

type Validator interface {
	ValidateStruct()
}

func CustomErrorMessage(err error) (map[string] string) {
	message := make(map[string] string)
	for _, err := range err.(validator.ValidationErrors) {
		condition := ""
		switch (err.Tag()) {
			case "min" :
				condition = " must be greater than " + err.Param() + " characters!"
				break
			case "max" : 
				condition = " must be less than " + err.Param() + " characters!"
				break
			case "required" : 
				condition = " is required!"
				break
			case "email" : 
				condition = " must be in the correct email format!"
				break
			case "alphanum" :
				condition = " can not contain special characters!"
				break
			case "numeric" :
				condition = " must be a number!"
				break
		}
		message[err.Field()+ strings.ToUpper(err.Tag())] = err.Field() + condition
	}
	return message
}
