package langs

func WritePhp(data *Data) {
	MakeLibraryDir("php")
	RunTemplate := ChooseTemplate("php")

	name := data.Pkg["name"].(string)

	RunTemplate("composer.json", "composer.json", data)
	MakeDir(name)

	RunTemplate("lib/Client.php", "Client.php", data)

	MakeDir("Exception")
	RunTemplate("lib/Exception/ExceptionInterface.php", "ExceptionInterface.php", data)
	RunTemplate("lib/Exception/ErrorException.php", "ErrorException.php", data)
	RunTemplate("lib/Exception/RuntimeException.php", "RuntimeException.php", data)
	MoveDir("..")

	MakeDir("HttpClient")
	RunTemplate("lib/HttpClient/HttpClient.php", "HttpClient.php", data)
	RunTemplate("lib/HttpClient/ResponseHandler.php", "ResponseHandler.php", data)
	RunTemplate("lib/HttpClient/RequestHandler.php", "RequestHandler.php", data)
	RunTemplate("lib/HttpClient/ErrorHandler.php", "ErrorHandler.php", data)
	RunTemplate("lib/HttpClient/AuthHandler.php", "AuthHandler.php", data)
	MoveDir("..")

	MakeDir("Api")
	MoveDir("..")
}
