<?php

namespace {{.Pkg.name}}\Api;

use {{.Pkg.name}}\HttpClient\HttpClient;

/*
 * {{index .Doc .Api.active.name "desc"}}
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
{{with $data := .}}{{range .Api.active.methods}}{{if (index $data.Api.class $data.Api.active.name . "method")}}
    /*
     *
     * '{{index $data.Api.class $data.Api.active.name . "path"}}' {{call $data.Fnc.upper (index $data.Api.class $data.Api.active.name . "method")}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index $data.Api.class $data.Api.active.name .) "body" true}})
    {

    }{{else}}
    /*
     *
     * '{{index $data.Api.class $data.Api.active.name . "path"}}' GET
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index $data.Api.class $data.Api.active.name .) "params" true}})
    {

    }
{{end}}{{end}}{{end}}
}
