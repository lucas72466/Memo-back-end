package user

// 实例

import "gorm.io/gorm"

type MySQLDBHandler struct {
	MySQLInst *gorm.DB
}

func (handler *MySQLDBHandler) FindUserByName(req *FindUserByNameRequest) *FindUserByNameResult {
	user := &User{}
	handler.MySQLInst.Debug().Where(&User{UserName: req.UserName}).Find(&user)

	res := &FindUserByNameResult{UserInfo: &Info{
		UserName: user.UserName,
		PassWord: user.PassWord,
	}}

	return res
}


