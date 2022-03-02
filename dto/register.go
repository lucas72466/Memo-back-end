package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type UserRegisterInput struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//从方法绑定参数

func (param *UserRegisterInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param); err != nil {
		return err
	}
	return nil
}

type UserRegisterOutput struct {
	Status int    `json:"status"`
	Note   string `json:"note"`
}
