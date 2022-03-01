package user

import "gorm.io/gorm"

// 依赖注入，创建数据库实例

var UDBHandler = InitDBHandler()

func InitMySQLInst() *gorm.DB {

}

func InitDBHandler() DBHandler {
	return &MySQLDBHandler{MySQLInst: InitMySQLInst()}
}
