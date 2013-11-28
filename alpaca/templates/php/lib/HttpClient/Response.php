<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Http\Message\Response;

class Response {

    //TODO: Use response header?
    public static function getContent(Response $response)
    {
        $body    = $response->getBody(true);
        $content = json_decode($body, true);

        if (JSON_ERROR_NONE !== json_last_error()) {
            return $body;
        }

        return $content;
    }

}
