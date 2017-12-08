package validations

import (
	
)

type BaseForm struct {
    Errors map[string]string
}
type BaseFormInterface interface {
	Validate()
}

