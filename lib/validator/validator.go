package validator

import (
	"encoding/json"
)

func IsValidJson(str string) bool {
	var js interface{}
	return json.Unmarshal([]byte(str), &js) == nil
}
