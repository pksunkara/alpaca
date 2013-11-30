<?php

namespace {{.Pkg.name}}\Api;

use {{.Pkg.name}}\HttpClient\HttpClient;

/*
 * {{index .Doc .Api.active.name "desc"}}
 *{{with $data := .}}{{call $data.Fnc.counter.start}}{{range .Api.active.args}}
 * @param ${{.}} {{index $data.Doc $data.Api.active.name "args" (call $data.Fnc.counter.value)}}{{end}}{{end}}
 * @param $client HttpClient Instance
 */
class {{call .Fnc.camelize .Api.active.name}}
{

{{range .Api.active.args}}    private ${{.}};
{{end}}    private $client;

    public function __construct({{call .Fnc.args.php (index .Api.class .Api.active.name) "args" false}}HttpClient $client)
    {
{{range .Api.active.args}}        $this->{{.}} = ${{.}};
{{end}}        $this->client = $client;
    }
{{with $data := .}}{{range .Api.active.methods}}
    /*
     * {{index $data.Doc $data.Api.active.name . "desc"}}{{if (index $data.Api.class $data.Api.active.name . "method")}}
     * '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (index $data.Api.class $data.Api.active.name . "method")}}
     *{{with $method := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Api.class $data.Api.active.name $method "body")}}
     * @param ${{.}} {{index $data.Doc $data.Api.active.name $method "body" (call $data.Fnc.counter.value)}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index $data.Api.class $data.Api.active.name .) "body" false}}array $options = array())
    {
        $body = array();{{range (index $data.Api.class $data.Api.active.name . "body")}}
        $body['{{.}}'] = ${{.}};{{end}}

        return $this->client->{{index $data.Api.class $data.Api.active.name . "method"}}{{else}}
     * '{{index $data.Api.class $data.Api.active.name . "path"}}' GET
     *{{with $method := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Api.class $data.Api.active.name $method "params")}}
     * @param ${{.}} {{index $data.Doc $data.Api.active.name $method "params" (call $data.Fnc.counter.value)}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index $data.Api.class $data.Api.active.name .) "params" false}}array $options = array())
    {
        $body = array();{{range (index $data.Api.class $data.Api.active.name . "params")}}
        $body['{{.}}'] = ${{.}};{{end}}

        return $this->client->get{{end}}("{{call $data.Fnc.urlr.php (index $data.Api.class $data.Api.active.name . "path") $data.Api.active.args}}", $body, $options);
    }
{{end}}{{end}}
}
