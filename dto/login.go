package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

// 定义input结构体
type UserLoginInput struct {
	UserName string `json:"user_name" binding:"required,max=20,min=5" customize_err_msg:"length of username should between 5-20"`
	Password string `json:"password" binding:"required,max=20,min=8,passwordValidate" customize_err_msg:"length of password should between 8-20 and at contain both upper case, lower case alphabet and number"`
}

// 绑定参数方法
func (param *UserLoginInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}

// 定义output结构体
type UserLoginOutput struct {
	status int    `json:"status"`
	Note   string `json:"note"`
	Token  string `json:"token"`
}
