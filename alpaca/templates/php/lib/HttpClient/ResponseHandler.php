<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Http\Message\Response;

class ResponseHandler {

    public static function getBody(Response $response)
    {
        $body = $response->getBody(true);
{{if .Api.response.formats.json}}
        if ($response->isContentType('json')) {
            $content = json_decode($body, true);

            if (JSON_ERROR_NONE === json_last_error()) {
                $body = $content;
            }
        }
{{end}}
        return $body;
    }

}
