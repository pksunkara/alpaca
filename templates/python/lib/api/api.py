class {{call .Fnc.camelize .Api.active.name}}(object):
    """{{index .Doc .Api.active.name "desc"}}

    Args:{{with $data := .}}{{range .Api.active.args}}
        {{.}}: {{index $data.Doc $data.Api.active.name "args" . "desc"}}{{end}}{{end}}
    """
{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}
    def __init__(self, {{call .Fnc.args.python .Api.active.args}}client):{{range .Api.active.args}}
        self.{{.}} = {{.}}{{end}}
        self.client = client
{{with $data := .}}{{range .Api.active.methods}}
    def {{call $data.Fnc.underscore .}}(self, {{call $data.Fnc.args.python (index (index $data.Api.class $data.Api.active.name .) "params")}}options={}):
        """{{index $data.Doc $data.Api.active.name . "desc"}}

        '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (or (index $data.Api.class $data.Api.active.name . "method") "get")}}

        Args:{{with $method := .}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}{{if .required}}
            {{.name}}: {{index $data.Doc $data.Api.active.name $method "params" .name "desc"}}{{end}}{{end}}{{end}}
        """
        body = options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}'] if '{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}' in options else {}{{range (index $data.Api.class $data.Api.active.name . "params")}}{{if .required}}
        body['{{.name}}'] = {{.name}}{{end}}{{end}}

        response = self.client.{{or (index $data.Api.class $data.Api.active.name . "method") "get"}}('{{call $data.Fnc.path.python (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}', body, options)

        return response
{{end}}{{end}}
