package conf

type StatusCode struct {
	Code int
	Msg  string
}

func NewStatusCode(code int, msg string) StatusCode {
	return StatusCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	Success            = NewStatusCode(0, "")
	InvalidParam       = NewStatusCode(1, "input param is invalid")
	InternalError      = NewStatusCode(2, "internal error")
	AuthenticationFail = NewStatusCode(3, "")

	RegisterSuccess   = NewStatusCode(100, "register successfully!")
	DuplicateUserName = NewStatusCode(101, "username has been used")
	UserNameNotFound  = NewStatusCode(102, "username doesn't exist")
	WrongPassword     = NewStatusCode(103, "username and password are mismatch")
	LoginSuccess      = NewStatusCode(104, "login successfully")

	CommentUploadSuccess = NewStatusCode(200, "comment has been upload successfully")

	PictureInvalid       = NewStatusCode(300, "upload picture is invalid, check format and size")
	PictureUploadSuccess = NewStatusCode(301, "picture upload successfully")
)
