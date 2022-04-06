package middleware

import (
	"Memo/conf"
	"Memo/public"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() func(c *gin.Context) {
	return jwtAuth
}

func jwtAuth(c *gin.Context) {
	tokenString, err := public.GetTokenFromContext(c)
	if err != nil {
		public.LogWithContext(c, public.ErrorLevel, err, nil)
		c.JSON(http.StatusUnauthorized, &public.DefaultResponse{
			StatusCode: conf.AuthenticationFail,
			Msg:        "token have not been set",
			Data:       nil,
		})
		c.Abort()
		return
	}

	_, claim, err := public.ParseUserToken(tokenString)
	if err != nil {
		public.LogWithContext(c, public.ErrorLevel, err, nil)
		c.JSON(http.StatusUnauthorized, &public.DefaultResponse{
			StatusCode: conf.AuthenticationFail,
			Msg:        "token is invalid",
			Data:       nil,
		})
		c.Abort()
		return
	}

	public.SetUserTokenInfoToContext(&public.UserTokenInfo{UserName: claim.UserName}, c)
}
