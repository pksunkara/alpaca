require "faraday"
require "json"
{{with $data := .}}{{range .Api.classes}}
require "{{call $data.Fnc.underscore $data.Pkg.Name}}/api/{{call $data.Fnc.underscore .}}"{{end}}{{end}}

module {{call .Fnc.camelize .Pkg.Name}}

  class Client

    def initialize(auth = {}, options = {})
      @http_client = {{call .Fnc.camelize .Pkg.Name}}::HttpClient::HttpClient.new auth, options
    end
{{with $data := .}}{{range .Api.classes}}
    # {{index $data.Doc . "desc"}}
    #{{with $class := .}}{{range (index $data.Api.class $class "args")}}
    # {{.}} - {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
    def {{call $data.Fnc.underscore .}}({{call $data.Fnc.args.ruby (index $data.Api.class . "args") true}})
      {{call $data.Fnc.camelize $data.Pkg.Name}}::Api::{{call $data.Fnc.camelize .}}.new {{call $data.Fnc.args.ruby (index $data.Api.class . "args")}}@http_client
    end
{{end}}{{end}}
  end

end
