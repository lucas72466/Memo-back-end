package public

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"unicode"
)

const (
	CustomizeErrMsgTagName = "customize_err_msg"
)

var customizeValidateFuncMap = map[string]validator.Func{
	"passwordValidate": passwordStrengthValidate,
}

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		InitCustomizeValidateFunc(v)
	}
}

func ParseValidatorErr(params interface{}, err error) error {
	validationErrs := err.(validator.ValidationErrors)
	fieldErrs := map[string]string{}
	for _, validationErr := range validationErrs {
		filedName := validationErr.Field()
		typeOfFiled := reflect.TypeOf(params)
		if typeOfFiled.Kind() == reflect.Ptr {
			typeOfFiled = typeOfFiled.Elem()
		}

		field, ok := typeOfFiled.FieldByName(filedName)
		if ok {
			filedJsonName := field.Tag.Get("json")
			errMsg := field.Tag.Get(CustomizeErrMsgTagName)
			fieldErrs[filedJsonName] = errMsg
		}
	}

	return errors.New(convertFieldErrs2Str(fieldErrs))
}

func convertFieldErrs2Str(fieldErrs map[string]string) string{
	bf := strings.Builder{}
	for fieldName, errInfo := range fieldErrs {
		bf.WriteString(fmt.Sprintf("%s:%s  ", fieldName, errInfo))
	}

	return bf.String()
}

func InitCustomizeValidateFunc(v *validator.Validate) {
	for key, customizeValidateFunc := range customizeValidateFuncMap {
		err := v.RegisterValidation(key, customizeValidateFunc)
		if err != nil {
			panic(err)
		}
	}
}

func passwordStrengthValidate(fl validator.FieldLevel) bool {
	if password, ok := fl.Field().Interface().(string); ok {
		var (
			existUpper  = false
			existLower  = false
			existNumber = false
		)

		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				existUpper = true
			case unicode.IsLower(char):
				existLower = true
			case unicode.IsNumber(char):
				existNumber = true
			}
		}

		if existUpper && existLower && existNumber {
			return true
		}
	}

	return false
}
