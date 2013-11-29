<?php

namespace {{.Pkg.name}}\HttpClient;

class RequestHandler {

    public static function createBody($body, array $options = array())
    {
        $type = isset($options['request_type']) ? $options['request_type'] : "{{.Api.request.formats.default}}";
{{if .Api.request.formats.json}}
        if ($type == 'json') {
            $body = ((count($body) === 0) ? null : json_encode($body, empty($body) ? JSON_FORCE_OBJECT : 0));
        }
{{end}}
        return $body;
    }

}
