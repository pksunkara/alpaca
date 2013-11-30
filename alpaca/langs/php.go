package langs

import (
	"bitbucket.org/pkg/inflect"
)

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

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/Api.php", inflect.Camelize(k)+".php", data)
		delete(data.Api, "active")
	}
}

func FunctionsPhp(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	urlr := fnc["urlr"].(map[string]interface{})

	args["php"] = ArgsFunctionMaker("$", ", ")
	urlr["php"] = UrlReplaceFunctionMaker("\".rawurlencode($$this->", ").\"")
}
