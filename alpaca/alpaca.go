package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"path/filepath"
	"strings"
)

var (
	LibraryRoot string
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

func WriteLibraries(directory string, opts *LanguageOptions) {
	var err error

	LibraryRoot, err = filepath.Abs(directory)
	HandleError(err)

	data := ReadData()
	ModifyData(data)

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
