package client

import (
	"encoding/json"
)

func isJson(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
