package public

import "regexp"

const (
	usernameRegexString = "^[a-z0-9A-Z]{5,15}$"
	passwordRegexString = "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$"
)

var (
	usernameRegex = regexp.MustCompile(usernameRegexString)
	//passwordRegex = regexp.MustCompile(passwordRegexString)
)
