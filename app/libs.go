package app

import "encoding/json"

// Pretty print object or variable
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
