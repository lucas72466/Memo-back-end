package conf

type StatusCode int

var (
	InvalidParam       StatusCode = 1
	InternalError      StatusCode = 2
	AuthenticationFail StatusCode = 3

	RegisterSuccess   StatusCode = 100
	DuplicateUserName StatusCode = 101
	UserNameNotFound  StatusCode = 102
	WrongPassword     StatusCode = 103
	LoginSuccess      StatusCode = 104
)

var ErrMsg = map[StatusCode]string{
	DuplicateUserName: "The username has been used",
	InternalError:     "Internal error",
	RegisterSuccess:   "Register successfully",
	UserNameNotFound:  "Username doesn't exist",
	WrongPassword:     "The password is mismatch",
}
