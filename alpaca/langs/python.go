package langs

import (
	"bitbucket.org/pkg/inflect"
)

func WritePython(data *Data) {
	MakeLibraryDir("python")
	RunTemplate := ChooseTemplate("python")

	name := data.Pkg["name"].(string)

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("setup.py", "setup.py", data)

	MakeDir(inflect.Underscore(name))
	RunTemplate("lib/__init__.py", "__init__.py", data)

	MakeDir("error")
	RunTemplate("lib/error/__init__.py", "__init__.py", data)
	RunTemplate("lib/error/client_error.py", "client_error.py", data)
	MoveDir("..")

	MakeDir("http_client")
	RunTemplate("lib/http_client/__init__.py", "__init__.py", data)
	RunTemplate("lib/http_client/auth_handler.py", "auth_handler.py", data)
	RunTemplate("lib/http_client/error_handler.py", "error_handler.py", data)
	RunTemplate("lib/http_client/request_handler.py", "request_handler.py", data)
	RunTemplate("lib/http_client/response_handler.py", "response_handler.py", data)
	MoveDir("..")

	MakeDir("api")

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/api.py", inflect.Underscore(k)+".py", data)
		delete(data.Api, "active")
	}
}

func FunctionsPython(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})

	args["python"] = ArgsFunctionMaker("", ", ")
	path["python"] = PathFunctionMaker("' + self.", " + '")
}
