{{define "boq"}}{{if (eq (or .Method "get") "get")}}query{{else}}body{{end}}{{end}}/**
 * {{(index .Doc .Active.Name).Desc}}{{with (index .Doc .Active.Name).Args}}
 *{{end}}{{with $data := .}}{{range .Active.Args}}
 * @param "{{.}}" {{(index ((index $data.Doc $data.Active.Name).Args) .).Desc}}{{end}}{{end}}
 */
var {{call .Fnc.camelize .Active.Name}} = function({{call .Fnc.args.node .Active.Args}}client) {
{{range .Active.Args}}  this.{{.}} = {{.}};
{{end}}  this.client = client;

  return this;
};
{{with $data := .}}{{range .Active.Functions}}
/**
 * {{(index ((index $data.Doc $data.Active.Name).Functions) .Name).Desc}}
 *
 * '{{.Path}}' {{call $data.Fnc.upper (or .Method "get")}}{{with .Params}}
 *{{end}}{{with $method := .}}{{range .Params}}{{if .Required}}
 * @param "{{.Name}}" {{(index ((index ((index $data.Doc $data.Active.Name).Functions) $method.Name).Params) .Name).Desc}}{{end}}{{end}}{{end}}
 */
{{call $data.Fnc.camelize $data.Active.Name}}.prototype.{{call $data.Fnc.camelizeDownFirst .Name}} = function ({{call $data.Fnc.args.node .Params}}options, callback) {
  if (typeof options === 'function') {
    callback = options;
    options = {};
  }

  var body = (options.{{template "boq" .}} ? options.{{template "boq" .}} : {});{{range .Params}}{{if .Required}}{{if (not .UrlUse)}}
  body['{{.Name}}'] = {{.Name}};{{end}}{{end}}{{end}}

  this.client.{{or .Method "get"}}('{{call $data.Fnc.path.node .Path $data.Active.Args .Params}}', body, options, function(err, response) {
    if (err) {
      return callback(err);
    }

    callback(null, response);
  });
};
{{end}}{{end}}
// Export module
module.exports = {{call .Fnc.camelize .Active.Name}};
