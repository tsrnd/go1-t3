package validations

import (
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
)

type RegisterForm struct {
	BaseForm
	Name string `form:"name" valid:"Required"`
	Email string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

func (form *RegisterForm) Validate() bool {
    valid := validation.Validation{}
	b, err := valid.Valid(form)
    if err != nil {
        beego.Error(err)
    }
    if !b {
		form.Errors = make(map[string]string)
        for _, err := range valid.Errors {
            form.Errors[err.Key] = err.Message
            beego.Debug(err.Key, err.Message)
        }
    }
    return b
}