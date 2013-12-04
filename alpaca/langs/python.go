package langs

func WritePython(data *Data) {
	MakeLibraryDir("python")
	RunTemplate := ChooseTemplate("python")

	//name := data.Pkg["name"].(string)

	RunTemplate("setup.py", "setup.py", data)
}

func FunctionsPython(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})

	args["python"] = ArgsFunctionMaker("", ", ")
	path["python"] = PathFunctionMaker("", "")
}
