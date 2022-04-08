package conf

import (
	"github.com/spf13/viper"
	"time"
)

// config for picture upload
var (
	PictureUploadMemoryLimit = 32 << 20
	PictureUploadKey         string
	PictureRelateBuildIDKey  string
	PictureStorageBucketName string
	PictureUploadTimeOut     time.Duration
)

func initPictureUploadConfig() {
	PictureUploadKey = viper.GetString("picture.uploadKey")
	PictureRelateBuildIDKey = viper.GetString("picture.relateBuildIDKey")
	PictureStorageBucketName = viper.GetString("picture.storageBucketName")
	PictureUploadTimeOut = time.Second * time.Duration(viper.GetInt("picture.uploadTimeOut"))
}

// config for token generate
var (
	DefaultUserTokenExpireTime time.Duration
	Issuer                     string
	SingedKey                  []byte
)

func initTokenConfig() {
	DefaultUserTokenExpireTime = time.Hour * time.Duration(viper.GetInt("token.expireTime"))
	Issuer = viper.GetString("token.issuer")
	SingedKey = []byte(viper.GetString("token.signedKey"))
}

// config for log
var (
	LogIDKey string
)

func initLogConfig() {
	LogIDKey = viper.GetString("log.logIDKey")
}

const (
	Decimal = 10
)
