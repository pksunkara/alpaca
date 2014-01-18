package alpaca

import (
	"bitbucket.org/pkg/inflect"
)

func WriteRuby(data *Data) {
	MakeLibraryDir("ruby")
	RunTemplate := ChooseTemplate("ruby")

	RunTemplate("gitignore", ".gitignore", data)
	RunTemplate("gemspec", data.Pkg.Package+".gemspec", data)
	RunTemplate("readme.md", "README.md", data)

	MakeDir("lib")
	RunTemplate("lib/name.rb", data.Pkg.Package+".rb", data)

	MakeDir(inflect.Underscore(data.Pkg.Name))
	RunTemplate("lib/client.rb", "client.rb", data)
	RunTemplate("lib/http_client.rb", "http_client.rb", data)
	RunTemplate("lib/error.rb", "error.rb", data)

	MakeDir("error")
	RunTemplate("lib/error/client_error.rb", "client_error.rb", data)
	MoveDir("..")

	MakeDir("http_client")
	RunTemplate("lib/http_client/auth_handler.rb", "auth_handler.rb", data)
	RunTemplate("lib/http_client/error_handler.rb", "error_handler.rb", data)
	RunTemplate("lib/http_client/request_handler.rb", "request_handler.rb", data)
	RunTemplate("lib/http_client/response.rb", "response.rb", data)
	RunTemplate("lib/http_client/response_handler.rb", "response_handler.rb", data)
	MoveDir("..")

	MakeDir("api")

	for k, v := range data.Api["class"].(map[string]interface{}) {
		data.Api["active"] = ActiveClassInfo(k, v)
		RunTemplate("lib/api/api.rb", inflect.Underscore(k)+".rb", data)
		delete(data.Api, "active")
	}
}

func FunctionsRuby(fnc map[string]interface{}) {
	args := fnc["args"].(map[string]interface{})
	path := fnc["path"].(map[string]interface{})
	prnt := fnc["prnt"].(map[string]interface{})

	args["ruby"] = ArgsFunctionMaker("", ", ")
	path["ruby"] = PathFunctionMaker("#{@", "}")
	prnt["ruby"] = PrntFunctionMaker(false, "  ", "\"", "\"", "[", "]", "{", "}", ":", " => ")
}

func CheckRuby(data *Data) error {
	return nil
}
