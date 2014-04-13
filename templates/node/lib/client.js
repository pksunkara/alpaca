/**
 * Main client for the module
 */
var Client = function({{if .Api.base_as_arg}}baseUrl, {{end}}auth, options) {
  this.httpClient = new (require('./http_client').HttpClient)({{if .Api.base_as_arg}}baseUrl, {{end}}auth, options);

  return this;
};
{{with $data := .}}{{range .Api.classes}}
/**
 * {{index $data.Doc . "desc"}}{{with (index $data.Api.class . "args")}}
 *{{end}}{{with $class := .}}{{range (index $data.Api.class $class "args")}}
 * @param "{{.}}" {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
 */
Client.prototype.{{call $data.Fnc.camelizeDownFirst .}} = function ({{call $data.Fnc.args.node (index $data.Api.class . "args") true}}) {
    return new (require('./api/{{call $data.Fnc.underscore .}}'))({{call $data.Fnc.args.node (index $data.Api.class . "args")}}this.httpClient);
};
{{end}}{{end}}
// Export module
module.exports = Client;
