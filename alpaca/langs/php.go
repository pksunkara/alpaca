package langs

import (
//"bitbucket.org/pkg/inflect"
)

func WritePhp(data *Data) {
	MakeLibraryDir("php")
	RunTemplate := ChooseTemplate("php")

	name := data.Pkg["name"].(string)

	RunTemplate("composer.json", "composer.json", data)
	MakeDir(name)

	MakeDir("Exception")
	RunTemplate("lib/Exception/ExceptionInterface.php", "ExceptionInterface.php", data)
	RunTemplate("lib/Exception/ErrorException.php", "ErrorException.php", data)
	RunTemplate("lib/Exception/RuntimeException.php", "RuntimeException.php", data)
	MoveDir("..")

	MakeDir("HttpClient")
	RunTemplate("lib/HttpClient/HttpClient.php", "HttpClient.php", data)
	RunTemplate("lib/HttpClient/Response.php", "Response.php", data)
	RunTemplate("lib/HttpClient/ErrorHandler.php", "ErrorHandler.php", data)
	RunTemplate("lib/HttpClient/AuthHandler.php", "AuthHandler.php", data)
	MoveDir("..")
}
