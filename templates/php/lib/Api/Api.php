<?php
{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}
namespace {{call .Fnc.camelize .Pkg.Name}}\Api;

use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\HttpClient;

/**
 * {{index .Doc .Api.active.name "desc"}}
 *{{with $data := .}}{{range .Api.active.args}}
 * @param ${{.}} {{index $data.Doc $data.Api.active.name "args" . "desc"}}{{end}}{{end}}
 */
class {{call .Fnc.camelize .Api.active.name}}
{

{{range .Api.active.args}}    private ${{.}};
{{end}}    private $client;

    public function __construct({{call .Fnc.args.php .Api.active.args}}HttpClient $client)
    {
{{range .Api.active.args}}        $this->{{.}} = ${{.}};
{{end}}        $this->client = $client;
    }
{{with $data := .}}{{range .Api.active.methods}}
    /**
     * {{index $data.Doc $data.Api.active.name . "desc"}}
     * '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (or (index $data.Api.class $data.Api.active.name . "method") "get")}}
     *{{with $method := .}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}{{if .required}}
     * @param ${{.name}} {{index $data.Doc $data.Api.active.name $method "params" .name "desc"}}{{end}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index (index $data.Api.class $data.Api.active.name .) "params")}}array $options = array())
    {
        $body = (isset($options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}']) ? $options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}'] : array());{{range (index $data.Api.class $data.Api.active.name . "params")}}{{if .required}}
        $body['{{.name}}'] = ${{.name}};{{end}}{{end}}

        $response = $this->client->{{or (index $data.Api.class $data.Api.active.name . "method") "get"}}('{{call $data.Fnc.path.php (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}', $body, $options);

        return $response;
    }
{{end}}{{end}}
}
