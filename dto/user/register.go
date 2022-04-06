package user

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type UserRegisterInput struct {
	UserName string `json:"user_name" binding:"required,max=20,min=5" customize_err_msg:"length of username should between 5-20"`
	Password string `json:"password" binding:"required,max=20,min=8,passwordValidate" customize_err_msg:"length of password should between 8-20 and at contain both upper case, lower case alphabet and number"`
}

//从方法绑定参数

func (param *UserRegisterInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param, true); err != nil {
		return err
	}
	return nil
}
