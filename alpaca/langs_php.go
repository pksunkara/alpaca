package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"errors"
)

func WritePhp(data *Data) {
	MakeLibraryDir("php")
	RunTemplate := ChooseTemplate("php")

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("composer.json", "composer.json", data)
	RunTemplate("readme.md", "README.md", data)

	MakeDir("lib")

	MakeDir(inflect.Camelize(data.Pkg.Name))
	RunTemplate("lib/Client.php", "Client.php", data)

	MakeDir("Exception")
	RunTemplate("lib/Exception/ExceptionInterface.php", "ExceptionInterface.php", data)
	RunTemplate("lib/Exception/ClientException.php", "ClientException.php", data)
	MoveDir("..")

	MakeDir("HttpClient")
	RunTemplate("lib/HttpClient/HttpClient.php", "HttpClient.php", data)
	RunTemplate("lib/HttpClient/AuthHandler.php", "AuthHandler.php", data)
	RunTemplate("lib/HttpClient/ErrorHandler.php", "ErrorHandler.php", data)
	RunTemplate("lib/HttpClient/RequestHandler.php", "RequestHandler.php", data)
	RunTemplate("lib/HttpClient/Response.php", "Response.php", data)
	RunTemplate("lib/HttpClient/ResponseHandler.php", "ResponseHandler.php", data)
	MoveDir("..")

	MakeDir("Api")

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/Api/Api.php", inflect.Camelize(k)+".php", data)
		delete(data.Api, "active")
	}
}

func FunctionsPhp(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})
	prnt := fnc["prnt"].(map[string]interface{})

	args["php"] = ArgsFunctionMaker("$", ", ")
	path["php"] = PathFunctionMaker("'.rawurlencode($$this->", ").'")
	prnt["php"] = PrntFunctionMaker(false, "    ", "\"", "\"", "array(", ")", "array(", ")", "'", "' => ")
}

func CheckPhp(data *Data) error {
	if data.Pkg.Php.Vendor == "" {
		return errors.New("php.vendor is needed in pkg.json for generating php library")
	}

	return nil
}
