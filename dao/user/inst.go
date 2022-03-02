package user

// 实例

import (
	"gorm.io/gorm"
)

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) FindUserByName(req *FindUserByNameRequest) *FindUserByNameResult {
	user := &User{}
	handler.MySQLInst.Debug().Where(&User{UserName: req.UserName}).First(user)

	res := &FindUserByNameResult{UserInfo: &Info{
		UserName: user.UserName,
		PassWord: user.PassWord},
	}

	return res
}

func (handler *MySQLDBHandler) CreateUser(req *CreateUserRequest) error {
	user := &User{UserName: req.UserName, PassWord: req.PassWord, ID: req.ID}

	if err := handler.MySQLInst.Debug().Create(&user).Error; err != nil {
		return err
	}

	return nil
}
