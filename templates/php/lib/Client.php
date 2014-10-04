<?php

namespace {{call .Fnc.camelize .Pkg.Name}};

use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\HttpClient;

class Client
{
    private $httpClient;

    public function __construct({{if .Api.BaseAsArg}}$baseUrl, {{end}}$auth = array(), array $options = array())
    {
        $this->httpClient = new HttpClient({{if .Api.BaseAsArg}}$baseUrl, {{end}}$auth, $options);
    }
{{with $data := .}}{{range .Api.Classes}}
    /**
     * {{(index $data.Doc .Name).Desc}}{{with .Args}}
     *{{end}}{{with $class := .}}{{range .Args}}
     * @param ${{.}} {{(index ((index $data.Doc $class.Name).Args) .).Desc}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .Name}}({{call $data.Fnc.args.php .Args true}})
    {
        return new Api\{{call $data.Fnc.camelize .Name}}({{call $data.Fnc.args.php .Args}}$this->httpClient);
    }
{{end}}{{end}}
}
