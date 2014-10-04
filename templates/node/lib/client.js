/**
 * Main client for the module
 */
var Client = function({{if .Api.BaseAsArg}}baseUrl, {{end}}auth, options) {
  this.httpClient = new (require('./http_client').HttpClient)({{if .Api.BaseAsArg}}baseUrl, {{end}}auth, options);

  return this;
};
{{with $data := .}}{{range .Api.Classes}}
/**
 * {{(index $data.Doc .Name).Desc}}{{with .Args}}
 *{{end}}{{with $class := .}}{{range .Args}}
 * @param "{{.}}" {{(index ((index $data.Doc $class.Name).Args) .).Desc}}{{end}}{{end}}
 */
Client.prototype.{{call $data.Fnc.camelizeDownFirst .Name}} = function ({{call $data.Fnc.args.node .Args true}}) {
    return new (require('./api/{{call $data.Fnc.underscore .Name}}'))({{call $data.Fnc.args.node .Args}}this.httpClient);
};
{{end}}{{end}}
// Export module
module.exports = Client;
