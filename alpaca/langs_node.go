package alpaca

import (
	"bitbucket.org/pkg/inflect"
)

func WriteNode(data *Data) {
	MakeLibraryDir("node")
	RunTemplate := ChooseTemplate("node")

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("package.json", "package.json", data)
	RunTemplate("readme.md", "README.md", data)

	MakeDir("lib")
	RunTemplate("lib/index.js", "index.js", data)

	MakeDir(inflect.CamelizeDownFirst(data.Pkg.Name))
	RunTemplate("lib/client.js", "client.js", data)

	MakeDir("error")
	RunTemplate("lib/error/index.js", "index.js", data)
	RunTemplate("lib/error/client_error.js", "client_error.js", data)
	MoveDir("..")

	MakeDir("http_client")
	RunTemplate("lib/http_client/index.js", "index.js", data)
	RunTemplate("lib/http_client/auth_handler.js", "auth_handler.js", data)
	RunTemplate("lib/http_client/error_handler.js", "error_handler.js", data)
	RunTemplate("lib/http_client/request_handler.js", "request_handler.js", data)
	RunTemplate("lib/http_client/response.js", "response.js", data)
	RunTemplate("lib/http_client/response_handler.js", "response_handler.js", data)
	MoveDir("..")

	MakeDir("api")

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/api/api.js", inflect.CamelizeDownFirst(k)+".js", data)
		delete(data.Api, "active")
	}
}

func FunctionsNode(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})
	prnt := fnc["prnt"].(map[string]interface{})

	args["node"] = ArgsFunctionMaker("", ", ")
	path["node"] = PathFunctionMaker("' + this.", " + '")
	prnt["node"] = PrntFunctionMaker(false, "  ", "\"", "\"", "[", "]", "{", "}", "", ": ")
}

func CheckNode(data *Data) error {
	return nil
}
