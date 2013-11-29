package langs

import (
	"bitbucket.org/pkg/inflect"
)

func WriteNode(data *Data) {
	MakeLibraryDir("node")
	RunTemplate := ChooseTemplate("node")

	name := data.Pkg["name"].(string)

	RunTemplate("package.json", "package.json", data)

	MakeDir("lib")
	RunTemplate("lib/index.js", "index.js", data)

	MakeDir(inflect.CamelizeDownFirst(name))
}

func FunctionsNode(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})

	args["node"] = ArgsTemplate("", ", ")
}
