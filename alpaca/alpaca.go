package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"encoding/json"
	"github.com/pksunkara/alpaca/alpaca/langs"
	"os"
	"path"
	"strings"
)

type LanguageOptions struct {
	Php    bool `long:"no-php" description:"Do not write php library"`
	Python bool `long:"no-python" description:"Do not write python library"`
	Ruby   bool `long:"no-ruby" description:"Do not write ruby library"`
	Node   bool `long:"no-node" description:"Do not write node library"`
}

func ReadData(directory string) *langs.Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON(directory+"/pkg.json", &pkg)
	ReadJSON(directory+"/api.json", &api)
	ReadJSON(directory+"/doc.json", &doc)

	return &langs.Data{pkg, api, doc, make(map[string]interface{})}
}

func WriteLibraries(directory string, opts *LanguageOptions) {
	data := ReadData(directory)
	ModifyData(data)

	langs.Init(HandleError, directory, "alpaca/templates")

	if !opts.Php {
		langs.WritePhp(data)
	}

	if !opts.Python {
		langs.WritePython(data)
	}

	if !opts.Ruby {
		langs.WriteRuby(data)
	}

	if !opts.Node {
		langs.WriteNode(data)
	}
}

func ModifyData(data *langs.Data) {
	data.Pkg["keywords"] = langs.ArrayInterfaceToString(data.Pkg["keywords"])
	data.Api["classes"] = langs.MapKeysToStringArray(data.Api["class"], []string{})

	data.Fnc["join"] = strings.Join
	data.Fnc["upper"] = strings.ToUpper

	data.Fnc["camelize"] = inflect.Camelize
	data.Fnc["camelizeDownFirst"] = inflect.CamelizeDownFirst
	data.Fnc["underscore"] = inflect.Underscore

	data.Fnc["counter"] = langs.CounterTracker()

	data.Fnc["args"] = make(map[string]interface{})
	data.Fnc["path"] = make(map[string]interface{})

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
