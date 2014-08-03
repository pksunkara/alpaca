{{define "boq"}}{{if (eq (or .Method "get") "get")}}query{{else}}body{{end}}{{end}}class {{call .Fnc.camelize .Active.Name}}(object):

    """{{(index .Doc .Active.Name).Desc}}{{with (index .Doc .Active.Name).Args}}

    Args:{{end}}{{with $data := .}}{{range .Active.Args}}
        {{.}}: {{(index ((index $data.Doc $data.Active.Name).Args) .).Desc}}{{end}}{{end}}
    """

    def __init__(self, {{call .Fnc.args.python .Active.Args}}client):{{range .Active.Args}}
        self.{{.}} = {{.}}{{end}}
        self.client = client
{{with $data := .}}{{range .Active.Functions}}
    def {{call $data.Fnc.underscore .Name}}(self, {{call $data.Fnc.args.python .Params}}options={}):
        """{{(index ((index $data.Doc $data.Active.Name).Functions) .Name).Desc}}

        '{{.Path}}' {{call $data.Fnc.upper (or .Method "get")}}{{with .Params}}

        Args:{{end}}{{with $method := .}}{{range .Params}}{{if .Required}}
            {{.Name}}: {{(index ((index ((index $data.Doc $data.Active.Name).Functions) $method.Name).Params) .Name).Desc}}{{end}}{{end}}{{end}}
        """
        body = options['{{template "boq" .}}'] if '{{template "boq" .}}' in options else {}{{range .Params}}{{if .Required}}{{if (not .UrlUse)}}
        body['{{.Name}}'] = {{.Name}}{{end}}{{end}}{{end}}

        response = self.client.{{or .Method "get"}}('{{call $data.Fnc.path.python .Path $data.Active.Args .Params}}', body, options)

        return response
{{end}}{{end}}
