package alpaca

import (
	"encoding/json"
	"os"
	"path"
)

func ReadData() *Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON("pkg.json", &pkg)
	ReadJSON("api.json", &api)
	ReadJSON("doc.json", &doc)

	return &Data{pkg, api, doc, make(map[string]interface{})}
}

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Join(LibraryRoot, name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	HandleError(os.RemoveAll(name))
	MakeDir(name)
}
