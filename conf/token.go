package conf

import (
	"github.com/spf13/viper"
	"time"
)

var (
	DefaultUserTokenExpireTime time.Duration
	Issuer                     string
	SingedKey                  []byte
)

func InitTokenConfig() {
	DefaultUserTokenExpireTime = time.Hour * time.Duration(viper.GetInt("token.expireTime"))
	Issuer = viper.GetString("token.issuer")
	SingedKey = []byte(viper.GetString("token.signedKey"))
}
