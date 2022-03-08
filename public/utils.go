package public

import "encoding/json"

func JsonString(input interface{}) string {
	content, err := json.Marshal(input)
	if err != nil {
		return ""
	}

	return string(content)
}
