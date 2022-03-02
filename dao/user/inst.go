package user

// 实例

import (
	"errors"
	"gorm.io/gorm"
)

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) FindUserByName(req *FindUserByNameRequest) (*FindUserByNameResult, error) {
	user := &User{}
	err := handler.MySQLInst.Debug().Where(&User{UserName: req.UserName}).First(user).Error
	if err != nil {
		return nil, err
	}

	res := &FindUserByNameResult{UserInfo: &Info{
		UserName: user.UserName,
		PassWord: user.PassWord},
	}

	return res, nil
}

func (handler *MySQLDBHandler) CreateUser(req *CreateUserRequest) error {
	if req == nil || req.UserInfo == nil {
		return errors.New("user info can not be empty")
	}
	user := &User{UserName: req.UserInfo.UserName, PassWord: req.UserInfo.PassWord}

	if err := handler.MySQLInst.Debug().Create(&user).Error; err != nil {
		return err
	}

	return nil
}
