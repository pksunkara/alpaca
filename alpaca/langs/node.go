package langs

func WriteNode(data interface{}) {
	MakeLibraryDir("node")

	RunTemplate("node/package.json", "package.json", data)
}
