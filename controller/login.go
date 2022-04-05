package controller

import (
	"Memo/conf"
	"Memo/dao/user"
	"Memo/dto"
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
	param := &dto.UserLoginInput{}

	if err := param.BindParam(c); err != nil {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.InvalidParam,
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
			StatusCode: conf.UserNameNotFound,
			Msg:        conf.ErrMsg[conf.UserNameNotFound],
			Data:       nil,
		}, err)
	}

	//3. 用户铭文的password用同样的方式加密，比对
	//4.如果密码错误，返回状态和文本；如果密码正确，返回状态文本和token
	if match, err := public.ComparePasswords(param.Password, res.UserInfo.PassWord); !match {
		public.ResponseError(c, &public.DefaultResponse{
			StatusCode: conf.WrongPassword,
			Msg:        conf.ErrMsg[conf.WrongPassword],
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
		StatusCode: conf.LoginSuccess,
		Msg:        conf.ErrMsg[conf.LoginSuccess],
		Data:       gin.H{"token": tokenString},
	})

}
