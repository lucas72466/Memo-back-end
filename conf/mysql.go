package conf

import "fmt"

var (
	defaultMySQLUserName = "admin"
	defaultMySQLPassword = "qaQLX2iy6jKjEDEyg4eZ"
	defaultMySQLDBName   = "Memo"
	defaultMySQLIP       = "memo-backend.ckasaaojz9uh.eu-west-2.rds.amazonaws.com"

	defaultMySQLLocalDSNTemplate = "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True"
)

var DefaultLocalMySQLDSN = generateDefaultLocalMySQLDSN()

func generateDefaultLocalMySQLDSN() string {
	return fmt.Sprintf(defaultMySQLLocalDSNTemplate, defaultMySQLUserName, defaultMySQLPassword, defaultMySQLIP, defaultMySQLDBName)
}
