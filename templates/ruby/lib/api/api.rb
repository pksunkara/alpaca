module {{call .Fnc.camelize .Pkg.Name}}
{{define "boq"}}{{if (eq (or .Method "get") "get")}}query{{else}}body{{end}}{{end}}
  module Api

    # {{(index .Doc .Active.Name).Desc}}{{with (index .Doc .Active.Name).Args}}
    #{{end}}{{with $data := .}}{{range .Active.Args}}
    # {{.}} - {{(index ((index $data.Doc $data.Active.Name).Args) .).Desc}}{{end}}{{end}}
    class {{call .Fnc.camelize .Active.Name}}

      def initialize({{call .Fnc.args.ruby .Active.Args}}client)
{{range .Active.Args}}        @{{.}} = {{.}}
{{end}}        @client = client
      end
{{with $data := .}}{{range .Active.Functions}}
      # {{(index ((index $data.Doc $data.Active.Name).Functions) .Name).Desc}}
      #
      # '{{.Path}}' {{call $data.Fnc.upper (or .Method "get")}}{{with .Params}}
      #{{end}}{{with $method := .}}{{range .Params}}{{if .Required}}
      # {{.Name}} - {{(index ((index ((index $data.Doc $data.Active.Name).Functions) $method.Name).Params) .Name).Desc}}{{end}}{{end}}{{end}}
      def {{call $data.Fnc.underscore .Name}}({{call $data.Fnc.args.ruby .Params}}options = {})
        body = options.fetch(:{{template "boq" .}}, {}){{range .Params}}{{if .Required}}{{if (not .UrlUse)}}
        body[:{{.Name}}] = {{.Name}}{{end}}{{end}}{{end}}

        @client.{{or .Method "get"}}("{{call $data.Fnc.path.ruby .Path $data.Active.Args .Params}}", body, options)
      end
{{end}}{{end}}
    end

  end

end
