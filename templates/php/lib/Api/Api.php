<?php
{{define "bodyorquery"}}{{if (eq (or (index . "method") "get") "get")}}query{{else}}body{{end}}{{end}}
namespace {{.Pkg.name}}\Api;

use {{.Pkg.name}}\HttpClient\HttpClient;

/**
 * {{index .Doc .Api.active.name "desc"}}
 *{{with $data := .}}{{range $index, $element := .Api.active.args}}
 * @param ${{.}} {{index $data.Doc $data.Api.active.name "args" $index}}{{end}}{{end}}
 * @param $client HttpClient Instance
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
     *{{with $method := .}}{{range $index, $element := (index $data.Api.class $data.Api.active.name $method "params")}}
     * @param ${{.}} {{index $data.Doc $data.Api.active.name $method "params" $index}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index (index $data.Api.class $data.Api.active.name .) "params")}}array $options = array())
    {
        $body = (isset($options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}']) ? $options['{{template "bodyorquery" (index $data.Api.class $data.Api.active.name .)}}'] : array());{{range (index $data.Api.class $data.Api.active.name . "params")}}
        $body['{{.}}'] = ${{.}};{{end}}

        $response = $this->client->{{or (index $data.Api.class $data.Api.active.name . "method") "get"}}('{{call $data.Fnc.path.php (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}', $body, $options);

        return array('body' => $response['body'], 'headers' => $response['headers']);
    }
{{end}}{{end}}
}
