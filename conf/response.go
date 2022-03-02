package conf

type ResponseCode int

var (
	RegisterSuccess   ResponseCode = 100
	DuplicateUserName ResponseCode = 200
	InvalidParam      ResponseCode = 300
	WrongPassword     ResponseCode = 400
	loginSuccess      ResponseCode = 500
	UserNameNotFound  ResponseCode = 600
)

var ErrMsg = map[ResponseCode]string{
	RegisterSuccess:   "Register successfully",
	DuplicateUserName: "The username has been used",
	InvalidParam:      "Username or Password is invalid",
	WrongPassword:     "The password is wrong",
	loginSuccess:      "Login successfully",
	UserNameNotFound:  "Username doesn't exist",
}
