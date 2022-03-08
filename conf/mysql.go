package conf

import "fmt"

var (
	defaultMySQLUserName = "root"
	defaultMySQLPassword = "Sunjn0503"
	defaultMySQLDBName   = "memo"

	defaultMySQLLocalDSNTemplate = "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True"
)

var DefaultLocalMySQLDSN = generateDefaultLocalMySQLDSN()

func generateDefaultLocalMySQLDSN() string {
	return fmt.Sprintf(defaultMySQLLocalDSNTemplate, defaultMySQLUserName, defaultMySQLPassword, defaultMySQLDBName)
}
