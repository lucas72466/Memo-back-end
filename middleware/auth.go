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
	claim, err := public.ParseTokenClaimFromContext(c)
	if err != nil {
		public.LogWithContext(c, public.ErrorLevel, err, nil)
		c.JSON(http.StatusUnauthorized, &public.DefaultResponse{
			StatusCode: conf.AuthenticationFail.Code,
			Msg:        "token is invalid",
			Data:       nil,
		})
		c.Abort()
	}

	public.SetUserTokenInfoToContext(&public.UserTokenInfo{UserName: claim.UserName}, c)
}


