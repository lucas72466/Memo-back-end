package user

import (
	"Memo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 依赖注入，创建数据库实例

var UDBHandler = InitDBHandler()

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
