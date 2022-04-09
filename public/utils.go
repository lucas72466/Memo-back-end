package public

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func JsonString(input interface{}) string {
	content, err := json.Marshal(input)
	if err != nil {
		return ""
	}

	return string(content)
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func CheckIsStringParamsAllEmpty(params ...string) bool {
	for _, param := range params {
		if param != "" {
			return false
		}
	}

	return true
}
