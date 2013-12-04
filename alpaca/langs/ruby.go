package langs

func WriteRuby(data *Data) {
	MakeLibraryDir("ruby")
	RunTemplate := ChooseTemplate("ruby")

	//name := data.Pkg["name"].(string)

	RunTemplate("gemspec", data.Pkg["package"].(string)+".gemspec", data)
}

func FunctionsRuby(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})

	args["ruby"] = ArgsFunctionMaker("", ", ")
	path["ruby"] = PathFunctionMaker("", "")
}
