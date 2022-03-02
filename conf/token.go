package conf

import "time"

var (
	DefaultUserTokenExpireTime = time.Hour * 24
	Issuer                     = "Memo"
	Subject                    = "user token"
	SingedKey                  = "Memo"
)
