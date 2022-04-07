package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	userDTO "Memo/dto/user"
	"Memo/public"
	"github.com/gin-gonic/gin"
	"log"
)

type UserLoginHandler struct{}

//路由登陆

func UserLoginRouteRegister(group *gin.RouterGroup) {
	handler := UserLoginHandler{}
	group.POST("/login", handler.UserLogin)
}
func (handler *UserLoginHandler) UserLogin(c *gin.Context) {

	//1.绑定校验参数
	param := &userDTO.UserLoginInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.InvalidParam.Code,
			Msg:        err.Error(),
			Data:       nil,
		}, err)
		return
	}

	//2. 查看用户名是否存在，根据username去数据库查询
	res, err := user.UDBHandler.FindUserByName(&user.FindUserByNameRequest{
		UserName: param.UserName,
	})

	if err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.UserNameNotFound.Code,
			Msg:        conf.UserNameNotFound.Msg,
			Data:       nil,
		}, err)
	}

	//3. 用户铭文的password用同样的方式加密，比对
	//4.如果密码错误，返回状态和文本；如果密码正确，返回状态文本和token
	if match, err := public.ComparePasswords(param.Password, res.UserInfo.PassWord); !match {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.WrongPassword.Code,
			Msg:        conf.WrongPassword.Msg,
			Data:       err,
		}, err)
		return
	}

	//密码正确
	tokenString, err := public.GenerateUserToken(param.UserName)
	if err != nil {
		log.Println(err)
		return
	}

	public.ResponseSuccess(c, &public.DefaultResponse{
		StatusCode: conf.LoginSuccess.Code,
		Msg:        conf.LoginSuccess.Msg,
		Data:       gin.H{"token": tokenString},
	})

}
