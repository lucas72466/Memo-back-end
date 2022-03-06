package dto

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
)

// 定义input结构体
type UserLoginInput struct {
	UserName string `json:"user_name" binding:"required,max=15,min=5"`
	Password string `json:"password" binding:"required,max=20,min=8" `
}

// 绑定参数方法
func (param *UserLoginInput) BindParam(c *gin.Context) error {
	if err := public.DefaultParamsBindAndValidate(c, param); err != nil {
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

// 定义登陆状态，是否成功并返回字段
type LoginStatus int

var (
	Success          LoginStatus = 100
	UserNameNotFound LoginStatus = 200
	InvalidUserName  LoginStatus = 300
	WrongPassword    LoginStatus = 400
	InvalidPassword  LoginStatus = 500
)

var LoginNote = map[LoginStatus]string{
	Success:          "Login success",
	UserNameNotFound: "Username doesn't exist",
	WrongPassword:    "The password is wrong",
	InvalidUserName:  "Username is invalid",
	InvalidPassword:  "Password is invalid",
}
