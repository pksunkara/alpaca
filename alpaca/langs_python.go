package alpaca

import (
	"bitbucket.org/pkg/inflect"
)

func WritePython(data *Data) {
	MakeLibraryDir("python")
	RunTemplate := ChooseTemplate("python")

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("setup.py", "setup.py", data)
	RunTemplate("readme.md", "README.md", data)

	MakeDir(inflect.Underscore(data.Pkg.Name))
	RunTemplate("lib/__init__.py", "__init__.py", data)
	RunTemplate("lib/client.py", "client.py", data)

	MakeDir("error")
	RunTemplate("lib/error/__init__.py", "__init__.py", data)
	RunTemplate("lib/error/client_error.py", "client_error.py", data)
	MoveDir("..")

	MakeDir("http_client")
	RunTemplate("lib/http_client/__init__.py", "__init__.py", data)
	RunTemplate("lib/http_client/auth_handler.py", "auth_handler.py", data)
	RunTemplate("lib/http_client/error_handler.py", "error_handler.py", data)
	RunTemplate("lib/http_client/request_handler.py", "request_handler.py", data)
	RunTemplate("lib/http_client/response.py", "response.py", data)
	RunTemplate("lib/http_client/response_handler.py", "response_handler.py", data)
	MoveDir("..")

	MakeDir("api")
	RunTemplate("lib/api/__init__.py", "__init__.py", data)

	for _, v := range data.Api.Classes {
		data.Active = &v
		RunTemplate("lib/api/api.py", inflect.Underscore(v.Name)+".py", data)
		data.Active = nil
	}
}

func FunctionsPython(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})
	prnt := fnc["prnt"].(map[string]interface{})

	args["python"] = ArgsFunctionMaker("", ", ")
	path["python"] = PathFunctionMaker("' + ", "self.", " + '")
	prnt["python"] = PrntFunctionMaker(true, "    ", "\"", "\"", "[", "]", "{", "}", "'", "': ")
}

func CheckPython(data *Data) error {
	return nil
}
