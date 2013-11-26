package alpaca

import (
	"encoding/json"
	"os"
	"path"
)

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Clean(name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}

func WriteJSON(name string, v interface{}) {
	file, err := os.Open(path.Clean(name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewEncoder(file).Encode(v))
}
