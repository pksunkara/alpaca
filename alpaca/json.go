package alpaca

import (
	"encoding/json"
	"os"
)

func ReadFile(name string, v interface{}) {
	file, err := os.Open(name)
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}

func WriteFile(name string, v interface{}) {
	file, err := os.Open(name)
	defer file.Close()
	HandleError(err)

	HandleError(json.NewEncoder(file).Encode(v))
}
