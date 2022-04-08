package conf

import "github.com/spf13/viper"

var (
	SecretKeyID string
	SecretKey   string
)

func initAWSConfig() {
	SecretKeyID = viper.GetString("IAM.secretKeyID")
	SecretKey = viper.GetString("IAM.secretKey")
}
