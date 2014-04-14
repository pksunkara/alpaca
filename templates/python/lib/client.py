from .http_client import HttpClient

# Assign all the api classes{{with $data := .}}{{range .Api.classes}}
from .api.{{call $data.Fnc.underscore .}} import {{call $data.Fnc.camelize .}}{{end}}{{end}}


class Client(object):

    def __init__(self, {{if .Api.base_as_arg}}base_url, {{end}}auth={}, options={}):
        self.http_client = HttpClient({{if .Api.base_as_arg}}base_url, {{end}}auth, options)
{{with $data := .}}{{range .Api.classes}}
    def {{call $data.Fnc.underscore .}}(self{{call $data.Fnc.args.python (index $data.Api.class . "args") true true}}):
        """{{index $data.Doc . "desc"}}{{with (index $data.Api.class . "args")}}

        Args:{{end}}{{with $class := .}}{{range (index $data.Api.class $class "args")}}
            {{.}}: {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
        """
        return {{call $data.Fnc.camelize .}}({{call $data.Fnc.args.python (index $data.Api.class . "args")}}self.http_client)
{{end}}{{end}}
