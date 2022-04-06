package public

import (
	"Memo/conf"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserTokenClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

const (
	tokenFieldName = "token"
)

// TODO token refresh automatically

func GenerateUserToken(username string) (string, error) {
	expireTime := time.Now().Add(conf.DefaultUserTokenExpireTime)
	claims := &UserTokenClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    conf.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(conf.SingedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetTokenFromContext(c *gin.Context) (string, error) {
	tokenString := c.GetHeader(tokenFieldName)

	if tokenString == "" {
		return "", errors.New("can not get token from header")
	}

	return tokenString, nil
}

func ParseUserToken(tokenString string) (*jwt.Token, *UserTokenClaims, error) {
	claims := &UserTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return conf.SingedKey, nil
	})
	return token, claims, err
}

type UserTokenInfo struct {
	UserName string
}

var userNameKey = "username"

func SetUserTokenInfoToContext(uTokenInfo *UserTokenInfo, c *gin.Context) {
	if uTokenInfo == nil {
		return
	}
	c.Set(userNameKey, uTokenInfo.UserName)
}

func GetUserTokenInfoFromContext(c *gin.Context) (*UserTokenInfo, error) {
	info := &UserTokenInfo{}
	if res, exists := c.Get(userNameKey); exists {
		info.UserName = res.(string)
	} else {
		return nil, errors.New("username do not exists")
	}

	return info, nil
}

func GetUserTokenInfoFromContextSilent(c *gin.Context) *UserTokenInfo {
	if info, err := GetUserTokenInfoFromContext(c); err != nil {
		return nil
	} else {
		return info
	}
}
