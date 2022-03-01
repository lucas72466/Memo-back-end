package user

// 实例

import "gorm.io/gorm"

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) FindUserByName(req *FindUserByNameRequest) *FindUserByNameResult {

}


