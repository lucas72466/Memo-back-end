package user

// 定义数据库结构

type User struct {
	ID         int64  `json:"id"`
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	CreateTime int64  `json:"create_time"  gorm:"autoCreateTime"`
	UpdateTime int64  `json:"update_time"  gorm:"autoUpdateTime:milli"`
	IsDelete   string `json:"is_delete"`
}

func (user *User) TableName() string {
	return "users"
}
