package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type UserRegisterInput struct {
	UserName string `json:"user_name" binding:"required,max=15,min=5"`
	Password string `json:"password" binding:"required,max=20,min=8"`
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
