package conf

type ResponseCode int

var (
	RegisterSuccess      ResponseCode = 100
	DuplicateUserName    ResponseCode = 200
	InvalidParam         ResponseCode = 300
	WrongPassword        ResponseCode = 400
	LoginSuccess         ResponseCode = 500
	UserNameNotFound     ResponseCode = 600
	InternalError        ResponseCode = 700
	InsufficientWord     ResponseCode = 800
	ExceedWordLimit      ResponseCode = 900
	PictureUploadFailure ResponseCode = 1000
	InvalidTitle         ResponseCode = 1100
	StoryUploadSuccess   ResponseCode = 1200
	CommentUploadSuccess ResponseCode = 1300
)

var ErrMsg = map[ResponseCode]string{
	RegisterSuccess:      "Register successfully",
	DuplicateUserName:    "The username has been used",
	InvalidParam:         "Username or Password is invalid",
	WrongPassword:        "The password is wrong",
	LoginSuccess:         "Login successfully",
	UserNameNotFound:     "Username doesn't exist",
	InternalError:        "Internal error",
	InsufficientWord:     "Content can not be empty",
	ExceedWordLimit:      "Comment Content should be less than 50 words",
	PictureUploadFailure: "The picture failed to upload",
	InvalidTitle:         "Title should be within 20 words and contains no special characters",
	StoryUploadSuccess:   "Story upload successfully",
	CommentUploadSuccess: "Comment upload successfully",
}
