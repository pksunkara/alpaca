module {{.Pkg.name}}
{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}
  module Api

    # {{index .Doc .Api.active.name "desc"}}
    #{{with $data := .}}{{call .Fnc.counter.start}}{{range .Api.active.args}}
    # {{.}} - {{index $data.Doc $data.Api.active.name "args" (call $data.Fnc.counter.value)}}{{end}}{{end}}
    class {{call .Fnc.camelize .Api.active.name}}

      def initialize({{call .Fnc.args.ruby (index .Api.class .Api.active.name) "args" false}}client)
{{range .Api.active.args}}        @{{.}} = {{.}}
{{end}}        @client = client
      end
{{with $data := .}}{{range .Api.active.methods}}
      # {{index $data.Doc $data.Api.active.name . "desc"}}
      # '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (or (index $data.Api.class $data.Api.active.name . "method") "get")}}
      #{{with $method := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}
      # {{.}} - {{index $data.Doc $data.Api.active.name $method "params" (call $data.Fnc.counter.value)}}{{end}}{{end}}
      def {{call $data.Fnc.underscore .}}({{call $data.Fnc.args.ruby (index $data.Api.class $data.Api.active.name .) "params" false}}options = {})
        body = options.has_key?(:{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}) ? options[:{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}] : {}{{range (index $data.Api.class $data.Api.active.name . "params")}}
        body[:{{.}}] = {{.}}{{end}}

        @client.{{or (index $data.Api.class $data.Api.active.name . "method") "get"}} "{{call $data.Fnc.path.ruby (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}", body, options
      end
{{end}}{{end}}
    end

  end

end
