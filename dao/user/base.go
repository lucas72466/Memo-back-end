package user

// 接口

type DBHandler interface {
	FindUserByName(req *FindUserByNameRequest) *FindUserByNameResult
	CreateUser(req *CreateUserRequest) error
}

type Info struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type FindUserByNameRequest struct {
	UserName string
}

type FindUserByNameResult struct {
	UserInfo *Info
}

type CreateUserRequest struct {
	UserName string
	PassWord string
	ID       int64
}
