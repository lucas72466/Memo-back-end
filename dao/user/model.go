package user

// 定义数据库结构

type User struct {
	ID         int64  `json:"id"`
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	Salt       string `json:"salt"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	IsDelete   string `json:"is_delete"`
}
