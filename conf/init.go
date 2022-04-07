package conf

import (
	"github.com/spf13/viper"
	"os"
)

const (
	configFolderName = "configs"
	configFileName   = "config"
	configFileType   = "yaml"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path + "/" + configFolderName)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	InitMySQLConfig()
	InitTokenConfig()
	InitAWSConfig()
}
