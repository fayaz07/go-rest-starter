package helpers

import "encoding/json"

func Pretty(i interface{}) string {
	str, _ := json.MarshalIndent(i, "", "\t")
	return string(str)
}
