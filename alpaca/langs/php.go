package langs

func WritePhp(data interface{}) {
	MakeLibraryDir("php")

	RunTemplate("php/composer.json", "composer.json", data)
}
