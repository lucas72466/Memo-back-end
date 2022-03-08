package user

// 定义数据库结构

type User struct {
	ID         int64  `json:"id"           gorm:"column:id; autoIncrement"`
	UserName   string `json:"user_name"    gorm:"column:user_name"`
	PassWord   string `json:"pass_word"    gorm:"column:password"`
	CreateTime int64  `json:"create_time"  gorm:"column:create_time; autoCreateTime:milli"`
	UpdateTime int64  `json:"update_time"  gorm:"column:update_time; autoUpdateTime:milli"`
	IsDelete   int    `json:"is_delete" gorm:"column:is_delete"`
}

func (user *User) TableName() string {
	return "user"
}
