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
    #{{with $class := .}}{{range (index $data.Api.class $class "args")}}
    # {{.}} - {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
    def {{call $data.Fnc.underscore .}}({{call $data.Fnc.args.ruby (index $data.Api.class . "args") true}})
      {{$data.Pkg.name}}::Api::{{call $data.Fnc.camelize .}}.new {{call $data.Fnc.args.ruby (index $data.Api.class . "args")}}@http_client
    end
{{end}}{{end}}
  end

end
