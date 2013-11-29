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
    public function {{call $data.Fnc.camelizeDownFirst .}} {
        return new Api\{{call $data.Fnc.camelize .}}($this->httpClient);
    }
{{end}}{{end}}
}
