require "faraday"
require "json"
{{with $data := .}}{{range .Api.classes}}
require "{{call $data.Fnc.underscore $data.Pkg.name}}/api/{{call $data.Fnc.underscore .}}"{{end}}{{end}}

module {{.Pkg.name}}

  class Client

    def initialize(auth = {}, options = {})
      @http_client = {{.Pkg.name}}::HttpClient::HttpClient.new auth, options
    end
{{with $data := .}}{{range .Api.classes}}
    # {{index $data.Doc . "desc"}}
    #{{with $class := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Doc $class "args")}}
    # {{index $data.Api.class $class "args" (call $data.Fnc.counter.value)}} - {{.}}{{end}}{{end}}
    def {{call $data.Fnc.underscore .}}({{call $data.Fnc.args.ruby (index $data.Api.class .) "args" true}})
      {{$data.Pkg.name}}::Api::{{call $data.Fnc.camelize .}}.new {{call $data.Fnc.args.ruby (index $data.Api.class .) "args" false}}@http_client
    end
{{end}}{{end}}
  end

end
