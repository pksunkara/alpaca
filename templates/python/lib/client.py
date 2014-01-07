from .http_client import HttpClient

# Assign all the api classes{{with $data := .}}{{range .Api.classes}}
from .api.{{call $data.Fnc.underscore .}} import {{call $data.Fnc.camelize .}}{{end}}{{end}}

class Client():

	def __init__(self, auth = {}, options = {}):
		self.http_client = HttpClient(auth, options)
{{with $data := .}}{{range .Api.classes}}
	# {{index $data.Doc . "desc"}}
	#{{with $class := .}}{{range (index $data.Api.class $class "args")}}
	# {{.}} - {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
	def {{call $data.Fnc.underscore .}}(self{{call $data.Fnc.args.python (index $data.Api.class . "args") true true}}):
		return {{call $data.Fnc.camelize .}}({{call $data.Fnc.args.python (index $data.Api.class . "args")}}self.http_client)
{{end}}{{end}}
