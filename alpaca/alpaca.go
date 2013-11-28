package alpaca

import (
	"./langs"
	"strings"
)

type Data struct {
	Pkg map[string]interface{}
	Api map[string]interface{}
	Doc map[string]interface{}
	Fnc map[string]interface{}
}

func ReadData(directory string) *Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON(directory+"/pkg.json", &pkg)
	ReadJSON(directory+"/api.json", &api)
	ReadJSON(directory+"/doc.json", &doc)

	return &Data{pkg, api, doc, make(map[string]interface{})}
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

func ModifyData(data *Data) {
	oldwords := data.Pkg["keywords"].([]interface{})
	keywords := make([]string, len(oldwords))

	for i, v := range oldwords {
		keywords[i] = v.(string)
	}
	data.Pkg["keywords"] = keywords

	data.Fnc["join"] = strings.Join
}
