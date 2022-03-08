package middleware

import (
	"Memo/public"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func JWTAuth() func(c *gin.Context) {
	return jwtAuth
}

func jwtAuth(c *gin.Context) {
	tokenString, err := public.GetTokenFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &public.DefaultResponse{
			ErrCode: 100,
			ErrMsg:  "token have not been set",
			Data:    nil,
		})
		c.Abort()
		return
	}

	_, claim, err := public.ParseUserToken(tokenString)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, &public.DefaultResponse{
			ErrCode: 100,
			ErrMsg:  "token is invalid",
			Data:    nil,
		})
		c.Abort()
		return
	}

	public.SetUserTokenInfoToContext(&public.UserTokenInfo{UserName: claim.UserName}, c)
}
