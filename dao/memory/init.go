package memory

import (
	"Memo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MDBHandler DBHandler

func InitMySQLInst() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.GetDefaultLocalMySQLDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func InitDBHandler() {
	MDBHandler = &MySQLDBHandler{MySQLInst: InitMySQLInst()}
}
