package conf

type StatusCode int

var (
	Empty              StatusCode = 0
	InvalidParam       StatusCode = 1
	InternalError      StatusCode = 2
	AuthenticationFail StatusCode = 3

	RegisterSuccess   StatusCode = 100
	DuplicateUserName StatusCode = 101
	UserNameNotFound  StatusCode = 102
	WrongPassword     StatusCode = 103
	LoginSuccess      StatusCode = 104

	CommentUploadSuccess StatusCode = 200
)

var StatusMsg = map[StatusCode]string{
	DuplicateUserName: "The username has been used",
	InternalError:     "Internal error",
	RegisterSuccess:   "Register successfully",
	UserNameNotFound:  "Username doesn't exist",
	WrongPassword:     "The password is mismatch",

	CommentUploadSuccess: "comment has been upload successfully",
}
