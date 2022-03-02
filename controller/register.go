package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	"Memo/dto"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type UserRegisterHandler struct {
}

//TODO:路由登陆

func UserRegisterRouteRegister(group *gin.RouterGroup) {
	userRegister := &UserLoginHandler{}
	group.POST("/login", UserLoginHandler{}.UserLogin)
}

func (handler *UserRegisterHandler) UserRegister(c *gin.Context) {

// 1. 绑定参数 + 参数校验

	param := &dto.UserRegisterInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InvalidParam,
			ErrMsg:  conf.ErrMsg[conf.InvalidParam],
			Data:    nil,
		}, err)
	}

// 2. 查询是否有重复的用户名
	input := user.FindUserByNameResult{
	UserName: param.UserName
	}
}

// 3. 在数据库生成记录+给password加密

// 4. 返回状态码
