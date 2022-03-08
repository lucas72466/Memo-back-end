package memory

import (
	"Memo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// TODO 复制粘贴的很乱 要改

var MDBHandler = InitDBHandler()

func InitMySQLInst() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.DefaultLocalMySQLDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("init mysql successful")

	return db
}

func InitDBHandler() DBHandler {
	return &MySQLDBHandler{MySQLInst: InitMySQLInst()}
}
