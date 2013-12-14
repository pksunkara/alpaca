from http_client import HttpClient

# Assign all the api classes{{with $data := .}}{{range .Api.classes}}
from api import {{call $data.Fnc.camelize .}}{{end}}{{end}}

class Client():

	def __init__(self, auth, options = {}):
		self.http_client = HttpClient(auth, options)
{{with $data := .}}{{range .Api.classes}}
	# {{index $data.Doc . "desc"}}
	#{{with $class := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Doc $class "args")}}
	# {{index $data.Api.class $class "args" (call $data.Fnc.counter.value)}} - {{.}}{{end}}{{end}}
	def {{call $data.Fnc.underscore .}}(self{{call $data.Fnc.args.python (index $data.Api.class . "args") true true}}):
		return {{call $data.Fnc.camelize .}}({{call $data.Fnc.args.python (index $data.Api.class . "args")}}self.http_client)
{{end}}{{end}}
