/**
 * Main client for the module
 */
var Client = function(auth, options) {
  this.httpClient = new (require('./http_client').HttpClient)(auth, options);

  return this;
};
{{with $data := .}}{{range .Api.classes}}
/**
 * {{index $data.Doc . "desc"}}
 *{{with $class := .}}{{range (index $data.Api.class $class "args")}}
 * @param "{{.}}" {{index $data.Doc $class "args" . "desc"}}{{end}}{{end}}
 */
Client.prototype.{{call $data.Fnc.camelizeDownFirst .}} = function ({{call $data.Fnc.args.node (index $data.Api.class . "args") true}}) {
    return new (require('./api/{{call $data.Fnc.camelizeDownFirst .}}'))({{call $data.Fnc.args.node (index $data.Api.class . "args")}}this.httpClient);
};
{{end}}{{end}}
// Export module
module.exports = Client;
