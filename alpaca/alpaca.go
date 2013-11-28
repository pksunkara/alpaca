package alpaca

import (
	"./langs"
	"encoding/json"
	"os"
	"path"
	"strings"
)

func ReadData(directory string) *langs.Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON(directory+"/pkg.json", &pkg)
	ReadJSON(directory+"/api.json", &api)
	ReadJSON(directory+"/doc.json", &doc)

	return &langs.Data{pkg, api, doc, make(map[string]interface{})}
}

func WriteLibraries(directory string) {
	data := ReadData(directory)
	ModifyData(data)

	langs.Init(HandleError, directory, "alpaca/templates")

	langs.WriteNode(data)
	langs.WritePhp(data)
	langs.WritePython(data)
	langs.WriteRuby(data)
}

func ModifyData(data *langs.Data) {
	oldwords := data.Pkg["keywords"].([]interface{})
	keywords := make([]string, len(oldwords))

	for i, v := range oldwords {
		keywords[i] = v.(string)
	}
	data.Pkg["keywords"] = keywords

	data.Fnc["join"] = strings.Join
}

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Clean(name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}
