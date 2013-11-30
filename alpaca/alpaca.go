package alpaca

import (
	"./langs"
	"bitbucket.org/pkg/inflect"
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
	data.Pkg["keywords"] = langs.ArrayInterfaceToString(data.Pkg["keywords"])
	data.Api["classes"] = langs.MapKeysToStringArray(data.Api["class"], []string{})

	data.Fnc["join"] = strings.Join
	data.Fnc["equal"] = strings.EqualFold
	data.Fnc["upper"] = strings.ToUpper

	data.Fnc["camelize"] = inflect.Camelize
	data.Fnc["camelizeDownFirst"] = inflect.CamelizeDownFirst
	data.Fnc["underscore"] = inflect.Underscore

	data.Fnc["counter"] = langs.CounterTracker()

	data.Fnc["args"] = make(map[string]interface{})
	data.Fnc["urlr"] = make(map[string]interface{})

	langs.FunctionsNode(data.Fnc)
	langs.FunctionsPhp(data.Fnc)
	langs.FunctionsPython(data.Fnc)
	langs.FunctionsRuby(data.Fnc)
}

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Clean(name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}
