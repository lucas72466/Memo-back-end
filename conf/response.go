package conf

type ResponseCode int

var (
	RegisterSuccess   ResponseCode = 100
	DuplicateUserName ResponseCode = 200
	InvalidParam      ResponseCode = 300
	WrongPassword     ResponseCode = 400
	LoginSuccess      ResponseCode = 500
	UserNameNotFound  ResponseCode = 600
	InternalError     ResponseCode = 700
)

var (
	StoryUploadSuccess ResponseCode = 800
	InvalidTitle       ResponseCode = 900
)

var ErrMsg = map[ResponseCode]string{
	RegisterSuccess:   "Register successfully",
	DuplicateUserName: "The username has been used",
	InvalidParam:      "Username or Password is invalid",
	WrongPassword:     "The password is wrong",
	LoginSuccess:      "Login successfully",
	UserNameNotFound:  "Username doesn't exist",
	InternalError:     "Internal error",

	StoryUploadSuccess: "Story upload successfully",
	InvalidTitle:       "Title is invalid",
}
