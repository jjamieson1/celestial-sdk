package utilities

import "encoding/json"

func RenderJson(f interface{}) ([]byte, error) {
	return json.MarshalIndent(f, "", "    ")
}
