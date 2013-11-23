package alpaca

import (
	"encoding/json"
	"os"
)

func ReadFile(name string) interface{} {
	var v interface{}

	file, err := os.Open(name)
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(&v))

	return v.(map[string]interface{})
}

func WriteFile(name string, v interface{}) {
	file, err := os.Open(name)
	defer file.Close()
	HandleError(err)

	HandleError(json.NewEncoder(file).Encode(v))
}
