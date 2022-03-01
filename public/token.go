package public

import (
	"Memo/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserTokenClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateUserToken(customizeClaim *UserTokenClaims) (string, error) {
	expireTime := time.Now().Add(conf.DefaultUserTokenExpireTime)
	claims := UserTokenClaims{
		UserName: customizeClaim.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    conf.Issuer,
			Subject:   conf.Subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(conf.SingedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
