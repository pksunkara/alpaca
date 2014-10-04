<?php
{{define "boq"}}{{if (eq (or .Method "get") "get")}}query{{else}}body{{end}}{{end}}
namespace {{call .Fnc.camelize .Pkg.Name}}\Api;

use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\HttpClient;

/**
 * {{(index .Doc .Active.Name).Desc}}{{with (index .Doc .Active.Name).Args}}
 *{{end}}{{with $data := .}}{{range .Active.Args}}
 * @param ${{.}} {{(index ((index $data.Doc $data.Active.Name).Args) .).Desc}}{{end}}{{end}}
 */
class {{call .Fnc.camelize .Active.Name}}
{

{{range .Active.Args}}    private ${{.}};
{{end}}    private $client;

    public function __construct({{call .Fnc.args.php .Active.Args}}HttpClient $client)
    {
{{range .Active.Args}}        $this->{{.}} = ${{.}};
{{end}}        $this->client = $client;
    }
{{with $data := .}}{{range .Active.Functions}}
    /**
     * {{(index ((index $data.Doc $data.Active.Name).Functions) .Name).Desc}}
     *
     * '{{.Path}}' {{call $data.Fnc.upper (or .Method "get")}}{{with .Params}}
     *{{end}}{{with $method := .}}{{range .Params}}{{if .Required}}
     * @param ${{.Name}} {{(index ((index ((index $data.Doc $data.Active.Name).Functions) $method.Name).Params) .Name).Desc}}{{end}}{{end}}{{end}}
     */
    public function {{call $data.Fnc.camelizeDownFirst .Name}}({{call $data.Fnc.args.php .Params}}array $options = array())
    {
        $body = (isset($options['{{template "boq" .}}']) ? $options['{{template "boq" .}}'] : array());{{range .Params}}{{if .Required}}{{if (not .UrlUse)}}
        $body['{{.Name}}'] = ${{.Name}};{{end}}{{end}}{{end}}

        $response = $this->client->{{or .Method "get"}}('{{call $data.Fnc.path.php .Path $data.Active.Args .Params}}', $body, $options);

        return $response;
    }
{{end}}{{end}}
}
