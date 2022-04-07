package user

import (
	"Memo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 依赖注入，创建数据库实例

var UDBHandler DBHandler

func InitMySQLInst() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.GetDefaultLocalMySQLDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func InitUserDBHandler() {
	UDBHandler = &MySQLDBHandler{MySQLInst: InitMySQLInst()}
}
