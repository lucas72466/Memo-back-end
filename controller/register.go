package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	userDTO "Memo/dto/user"
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
	param := &userDTO.UserRegisterInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.InvalidParam.Code,
			Msg:        err.Error(),
			Data:       nil,
		}, err)
		return
	}

	// 2. 查询是否有重复的用户名
	if exist := duplicateUser(param.UserName); exist {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.DuplicateUserName.Code,
			Msg:        conf.DuplicateUserName.Msg,
			Data:       nil,
		}, errors.New(""))
		return
	}

	// 3. 在数据库生成记录+给password加密
	hashedPassword, err := public.GenerateHashedPassword(param.Password)
	if err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.InternalError.Code,
			Msg:        conf.InternalError.Msg,
			Data:       nil,
		}, err)
		return
	}

	userInfo := &user.Info{
		UserName: param.UserName,
		PassWord: hashedPassword,
	}
	if err := user.UDBHandler.CreateUser(&user.CreateUserRequest{
		UserInfo: userInfo,
	}); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.InternalError.Code,
			Msg:        conf.InternalError.Msg,
			Data:       nil,
		}, err)
		return
	}

	// 4. 返回状态码
	public.ResponseSuccess(c, &public.DefaultResponse{
		StatusCode: conf.RegisterSuccess.Code,
		Msg:        conf.RegisterSuccess.Msg,
		Data:       nil,
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
