{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}/**
 * {{index .Doc .Api.active.name "desc"}}
 *{{with $data := .}}{{call .Fnc.counter.start}}{{range .Api.active.args}}
 * @param {{.}} {{index $data.Doc $data.Api.active.name "args" (call $data.Fnc.counter.value)}}{{end}}{{end}}
 */
var {{call .Fnc.camelize .Api.active.name}} = function({{call .Fnc.args.node .Api.active.args}}client) {
{{range .Api.active.args}}  this.{{.}} = {{.}};
{{end}}  this.client = client;

  return this;
};
{{with $data := .}}{{range .Api.active.methods}}
/**
 * {{index $data.Doc $data.Api.active.name . "desc"}}
 * '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (or (index $data.Api.class $data.Api.active.name . "method") "get")}}
 *{{with $method := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}
 * @param "{{.}}" {{index $data.Doc $data.Api.active.name $method "params" (call $data.Fnc.counter.value)}}{{end}}{{end}}
 */
{{call $data.Fnc.camelize $data.Api.active.name}}.prototype.{{call $data.Fnc.camelizeDownFirst .}} = function ({{call $data.Fnc.args.node (index (index $data.Api.class $data.Api.active.name .) "params")}}options, callback) {
  if (typeof options == 'function') {
    callback = options;
    options = {};
  }

  var body = (options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}'] ? options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}'] : {});{{range (index $data.Api.class $data.Api.active.name . "params")}}
  body['{{.}}'] = {{.}};{{end}}

  this.client.{{or (index $data.Api.class $data.Api.active.name . "method") "get"}}('{{call $data.Fnc.path.node (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}', body, options, function(err, body, status, headers) {
    if (err) {
      return callback(err);
    }

    callback(null, body, headers);
  });
};
{{end}}{{end}}
// Export module
module.exports = {{call .Fnc.camelize .Api.active.name}}
