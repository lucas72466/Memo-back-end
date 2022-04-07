package conf

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const (
	defaultConfigFolderName = "configs"
	defaultConfigFileName   = "config"
	defaultConfigFileType   = "yaml"
)

func InitConfig() {
	initPFlag()

	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	InitMySQLConfig()
	InitTokenConfig()
	InitAWSConfig()
	InitLogConfig()
}

func initPFlag() {
	pflag.String("mode", "debug", "gin running mode")
	pflag.String("port", "8080", "server running port")
	pflag.String("config", getDefaultConfigPath(), "config file path")

	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
}

func getDefaultConfigPath() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s/%s.%s", workPath, defaultConfigFolderName, defaultConfigFileName, defaultConfigFileType)
}
