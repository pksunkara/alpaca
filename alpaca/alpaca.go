package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"encoding/json"
	"os"
	"path"
	"strings"
)

type Data struct {
	Pkg map[string]interface{}
	Api map[string]interface{}
	Doc map[string]interface{}
	Fnc map[string]interface{}
}

type LanguageOptions struct {
	Php    bool `long:"no-php" description:"Do not write php library"`
	Python bool `long:"no-python" description:"Do not write python library"`
	Ruby   bool `long:"no-ruby" description:"Do not write ruby library"`
	Node   bool `long:"no-node" description:"Do not write node library"`
}

func ReadData(directory string) *Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON(directory+"/pkg.json", &pkg)
	ReadJSON(directory+"/api.json", &api)
	ReadJSON(directory+"/doc.json", &doc)

	return &Data{pkg, api, doc, make(map[string]interface{})}
}

func WriteLibraries(directory string, opts *LanguageOptions) {
	data := ReadData(directory)
	ModifyData(data)

	TemplateInit(directory, "templates")

	if !opts.Php {
		WritePhp(data)
	}

	if !opts.Python {
		WritePython(data)
	}

	if !opts.Ruby {
		WriteRuby(data)
	}

	if !opts.Node {
		WriteNode(data)
	}
}

func ModifyData(data *Data) {
	data.Pkg["keywords"] = ArrayInterfaceToString(data.Pkg["keywords"])
	data.Api["classes"] = MapKeysToStringArray(data.Api["class"], []string{})

	data.Fnc["join"] = strings.Join
	data.Fnc["upper"] = strings.ToUpper

	data.Fnc["camelize"] = inflect.Camelize
	data.Fnc["camelizeDownFirst"] = inflect.CamelizeDownFirst
	data.Fnc["underscore"] = inflect.Underscore

	data.Fnc["counter"] = CounterTracker()

	data.Fnc["args"] = make(map[string]interface{})
	data.Fnc["path"] = make(map[string]interface{})

	FunctionsNode(data.Fnc)
	FunctionsPhp(data.Fnc)
	FunctionsPython(data.Fnc)
	FunctionsRuby(data.Fnc)
}

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Clean(name))
	defer file.Close()
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	HandleError(os.RemoveAll(name))
	MakeDir(name)
}
