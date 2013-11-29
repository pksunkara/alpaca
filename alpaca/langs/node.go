package langs

import (
	"bitbucket.org/pkg/inflect"
)

func WriteNode(data *Data) {
	InflectionsNode(data)
	MakeLibraryDir("node")
	RunTemplate := ChooseTemplate("node")

	name := data.Pkg["name"].(string)

	RunTemplate("package.json", "package.json", data)

	MakeDir("lib")
	RunTemplate("lib/index.js", "index.js", data)

	MakeDir(inflect.CamelizeDownFirst(name))
}

func InflectionsNode(data *Data) {
	data.Fnc["classify"] = inflect.CamelizeDownFirst
}
