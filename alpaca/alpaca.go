package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version = "0.2.0"
)

var (
	LibraryRoot string
	FormatList  []string
)

type Data struct {
	Pkg PkgStruct
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

func LoadLibraryPath(directory string) {
	var err error

	LibraryRoot, err = filepath.Abs(directory)
	HandleError(err)
}

func ConvertFormat(format string) {
	acceptable := false

	FormatList = []string{"blueprint"}

	for _, v := range FormatList {
		if v == format {
			acceptable = true
		}
	}

	if !acceptable {
		fmt.Println("The given format is not allowed. Please choose one from the following:\n")
		fmt.Println(strings.Join(FormatList, ", ") + "\n")
		os.Exit(0)
	}
}

func WriteLibraries(opts *LanguageOptions) {
	data := ReadData()
	ModifyData(data)

	if !opts.Php {
		HandleError(CheckPhp(data))
		FunctionsPhp(data.Fnc)
		WritePhp(data)
	}

	if !opts.Python {
		HandleError(CheckPython(data))
		FunctionsPython(data.Fnc)
		WritePython(data)
	}

	if !opts.Ruby {
		HandleError(CheckRuby(data))
		FunctionsRuby(data.Fnc)
		WriteRuby(data)
	}

	if !opts.Node {
		HandleError(CheckNode(data))
		FunctionsNode(data.Fnc)
		WriteNode(data)
	}
}

func ReadData() *Data {
	var pkg PkgStruct
	var api, doc map[string]interface{}

	ReadJSON("pkg.json", &pkg)
	ReadJSON("api.json", &api)
	ReadJSON("doc.json", &doc)

	return &Data{pkg, api, doc, make(map[string]interface{})}
}

func ModifyData(data *Data) {
	data.Api["alpaca_version"] = Version

	data.Api["classes"] = MapKeysToStringArray(data.Api["class"], []string{})

	data.Fnc["join"] = strings.Join
	data.Fnc["upper"] = strings.ToUpper

	data.Fnc["camelize"] = inflect.Camelize
	data.Fnc["camelizeDownFirst"] = inflect.CamelizeDownFirst
	data.Fnc["underscore"] = inflect.Underscore

	data.Fnc["methods"] = MethodList

	data.Fnc["args"] = make(map[string]interface{})
	data.Fnc["path"] = make(map[string]interface{})
	data.Fnc["prnt"] = make(map[string]interface{})
}
