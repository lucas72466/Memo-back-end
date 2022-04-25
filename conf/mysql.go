package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	defaultMySQLUserName string
	defaultMySQLPassword string
	defaultMySQLDBName   string
	defaultMySQLHost     string

	defaultMySQLLocalDSNTemplate = "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True"
)

var (
	DefaultCommentPageSize      int
	DefaultCommentPageSizeLimit int
)

func initMySQLConfig() {
	defaultMySQLUserName = viper.GetString("database.MySQLUserName")
	defaultMySQLPassword = viper.GetString("database.MySQLPassword")
	defaultMySQLDBName = viper.GetString("database.MySQLDBName")
	defaultMySQLHost = viper.GetString("database.MySQLDBHost")

	initPageConfig()
}

func GetDefaultLocalMySQLDSN() string {
	return fmt.Sprintf(defaultMySQLLocalDSNTemplate, defaultMySQLUserName, defaultMySQLPassword, defaultMySQLHost, defaultMySQLDBName)
}

func initPageConfig() {
	DefaultCommentPageSize = viper.GetInt("database.defaultCommentPageSize")
	DefaultCommentPageSizeLimit = viper.GetInt("database.defaultCommentPageSizeLimit")
}
