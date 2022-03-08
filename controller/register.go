package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	"Memo/dto"
	"Memo/public"
	"errors"
	"github.com/gin-gonic/gin"
)

type UserRegisterHandler struct {
}

func UserRegisterRouteRegister(group *gin.RouterGroup) {
	handler := UserRegisterHandler{}
	group.POST("/register", handler.UserRegister)
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
		return
	}

	// 2. 查询是否有重复的用户名
	if exist := duplicateUser(param.UserName); exist {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.DuplicateUserName,
			ErrMsg:  conf.ErrMsg[conf.DuplicateUserName],
			Data:    nil,
		}, errors.New(""))
		return
	}

	// 3. 在数据库生成记录+给password加密
	hashedPassword, err := public.GenerateHashedPassword(param.Password)
	if err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InternalError,
			ErrMsg:  conf.ErrMsg[conf.InternalError],
			Data:    nil,
		}, err)
	}

	userInfo := &user.Info{
		UserName: param.UserName,
		PassWord: hashedPassword,
	}
	//
	if err := user.UDBHandler.CreateUser(&user.CreateUserRequest{
		UserInfo: userInfo,
	}); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InternalError,
			ErrMsg:  conf.ErrMsg[conf.InternalError],
			Data:    nil,
		}, err)

	}

	// 4. 返回状态码
	public.ResponseSuccess(c, &public.DefaultResponse{
		ErrCode: conf.RegisterSuccess,
		ErrMsg:  conf.ErrMsg[conf.RegisterSuccess],
		Data:    nil,
	})
}

func duplicateUser(username string) bool {
	_, err := user.UDBHandler.FindUserByName(&user.FindUserByNameRequest{
		UserName: username,
	})
	if err == nil {
		return true
	}

	return false
}
