<?php

namespace {{.Pkg.name}};

use {{.Pkg.name}}\HttpClient\HttpClient;

class Client
{
    private $options;
    private $httpClient;

    public function __construct($auth, array $options = array())
    {
        $this->options = $options;
        $this->httpClient = new HttpClient($auth, $options);
    }
{{with $data := .}}{{range .Api.classes}}
    /*
     * {{index $data.Doc . "desc"}}
     *{{with $class := .}}{{call $data.Fnc.counter.start}}{{range (index $data.Doc . "args")}}
     * @param ${{index $data.Api.class $class "args" (call $data.Fnc.counter.value)}} {{.}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .}}({{call $data.Fnc.args.php (index $data.Api.class .) true}}) {
        return new Api\{{call $data.Fnc.camelize .}}({{call $data.Fnc.args.php (index $data.Api.class .) false}}$this->httpClient);
    }
{{end}}{{end}}
}
