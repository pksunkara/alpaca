module {{call .Fnc.camelize .Pkg.Name}}
{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}
  module Api

    # {{index .Doc .Api.active.name "desc"}}
    #{{with $data := .}}{{range .Api.active.args}}
    # {{.}} - {{index $data.Doc $data.Api.active.name "args" . "desc"}}{{end}}{{end}}
    class {{call .Fnc.camelize .Api.active.name}}

      def initialize({{call .Fnc.args.ruby .Api.active.args}}client)
{{range .Api.active.args}}        @{{.}} = {{.}}
{{end}}        @client = client
      end
{{with $data := .}}{{range .Api.active.methods}}
      # {{index $data.Doc $data.Api.active.name . "desc"}}
      # '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (or (index $data.Api.class $data.Api.active.name . "method") "get")}}
      #{{with $method := .}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}{{if .required}}
      # {{.name}} - {{index $data.Doc $data.Api.active.name $method "params" .name "desc"}}{{end}}{{end}}{{end}}
      def {{call $data.Fnc.underscore .}}({{call $data.Fnc.args.ruby (index (index $data.Api.class $data.Api.active.name .) "params")}}options = {})
        body = options.has_key?(:{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}) ? options[:{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}] : {}{{range (index $data.Api.class $data.Api.active.name . "params")}}{{if .required}}
        body[:{{.name}}] = {{.name}}{{end}}{{end}}

        response = @client.{{or (index $data.Api.class $data.Api.active.name . "method") "get"}} "{{call $data.Fnc.path.ruby (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}", body, options

        return response
      end
{{end}}{{end}}
    end

  end

end
