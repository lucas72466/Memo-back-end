package public

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	usernameRegexString = "^[0-9A-Za-z]{5,15}$"

	passwordRegexString = "^[(0-9)+(a-zA-Z)+]{8,20}$"
)

var (
	usernameRegex = regexp.MustCompile(usernameRegexString)
	//	passwordRegex = regexp.MustCompile(passwordRegexString)
)

var customizeValidateFuncMap = map[string]validator.Func{
	"passwordValidate": passwordValidate,
}

func InitCustomizeValidateFunc() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for key, customizeValidateFunc := range customizeValidateFuncMap {
			err := v.RegisterValidation(key, customizeValidateFunc)
			if err != nil {
				panic(err)
			}
		}
	}
}

func passwordValidate(fl validator.FieldLevel) bool {
	if password, ok := fl.Field().Interface().(string); ok {

	}
}
