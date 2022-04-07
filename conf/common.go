package conf

import (
	"github.com/spf13/viper"
	"time"
)

// config for picture upload
const (
	PictureUploadMemoryLimit = 32 << 20
	PictureUploadKey         = "file"
	PictureRelateBuildIDKey  = "building_id"
	PictureStorageBucketName = "memo-backend"
	PictureUploadTimeOut     = time.Second * 5
)

// config for token generate
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

// config for log
var (
	LogIDKey string
)

func InitLogConfig() {
	LogIDKey = viper.GetString("log.logIDKey")
}

const (
	Decimal = 10
)
