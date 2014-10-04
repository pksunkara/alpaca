from .http_client import HttpClient

# Assign all the api classes{{with $data := .}}{{range .Api.Classes}}
from .api.{{call $data.Fnc.underscore .Name}} import {{call $data.Fnc.camelize .Name}}{{end}}{{end}}


class Client(object):

    def __init__(self, {{if .Api.BaseAsArg}}base_url, {{end}}auth={}, options={}):
        self.http_client = HttpClient({{if .Api.BaseAsArg}}base_url, {{end}}auth, options)
{{with $data := .}}{{range .Api.Classes}}
    def {{call $data.Fnc.underscore .Name}}(self{{call $data.Fnc.args.python .Args true true}}):
        """{{(index $data.Doc .Name).Desc}}{{with .Args}}

        Args:{{end}}{{with $class := .}}{{range .Args}}
            {{.}}: {{(index ((index $data.Doc $class.Name).Args) .).Desc}}{{end}}{{end}}
        """
        return {{call $data.Fnc.camelize .Name}}({{call $data.Fnc.args.python .Args}}self.http_client)
{{end}}{{end}}
