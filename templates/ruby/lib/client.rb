require "faraday"
require "json"
{{with $data := .}}{{range .Api.Classes}}
require "{{call $data.Fnc.underscore $data.Pkg.Name}}/api/{{call $data.Fnc.underscore .Name}}"{{end}}{{end}}

module {{call .Fnc.camelize .Pkg.Name}}

  class Client

    def initialize({{if .Api.BaseAsArg}}base_url, {{end}}auth = {}, options = {})
      @http_client = {{call .Fnc.camelize .Pkg.Name}}::HttpClient::HttpClient.new({{if .Api.BaseAsArg}}base_url, {{end}}auth, options)
    end
{{with $data := .}}{{range .Api.Classes}}
    # {{(index $data.Doc .Name).Desc}}{{with .Args}}
    #{{end}}{{with $class := .}}{{range .Args}}
    # {{.}} - {{(index ((index $data.Doc $class.Name).Args) .).Desc}}{{end}}{{end}}
    def {{call $data.Fnc.underscore .Name}}({{call $data.Fnc.args.ruby .Args true}})
      {{call $data.Fnc.camelize $data.Pkg.Name}}::Api::{{call $data.Fnc.camelize .Name}}.new({{call $data.Fnc.args.ruby .Args}}@http_client)
    end
{{end}}{{end}}
  end

end
