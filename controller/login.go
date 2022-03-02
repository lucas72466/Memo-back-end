package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	"Memo/dto"
	"Memo/public"
	"github.com/gin-gonic/gin"
)

type UserLoginHandler struct{}

//路由登陆

func UserLoginRouteRegister(group *gin.RouterGroup) {
	userLogin := &UserLoginHandler{}
	group.POST("/login", UserLoginHandler{}.UserLogin)
}
func (handler *UserLoginHandler) UserLogin(c *gin.Context) {

	//1.绑定校验参数
	param := &dto.UserLoginInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: conf.InvalidParam,
			ErrMsg:  conf.ErrMsg[conf.InvalidParam],
			Data:    nil,
		}, err)
	}

	//2.1 查看用户名是否存在，根据username去数据库查询
	res := user.UDBHandler.FindUserByName(&user.FindUserByNameRequest{
		UserName: param.UserName,
	})

	if bool := public.ComparePasswords(res.UserInfo.PassWord, param.Password); bool == true{
		public.ResponseError(c, &public.DefaultResponse{
			ErrCode: 0,
			ErrMsg:  "",
			Data:    nil,
		}),
	}



	//3.用户铭文的password用同样的方式加密，比对

	//4.如果密码错误，返回状态和文本；如果密码正确，返回状态文本和token

}
