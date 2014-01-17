<?php

namespace {{call .Fnc.camelize .Pkg.Name}}\HttpClient;

use Guzzle\Http\Message\RequestInterface;

/**
 * RequestHandler takes care of encoding the request body into format given by options
 */
class RequestHandler {

    public static function setBody(RequestInterface $request, $body, $options)
    {
        $type = isset($options['request_type']) ? $options['request_type'] : '{{or .Api.request.formats.default "raw"}}';
        $header = null;
{{if .Api.request.formats.json}}
        // Encoding request body into JSON format
        if ($type == 'json') {
            $body = ((count($body) === 0) ? '{}' : json_encode($body, empty($body) ? JSON_FORCE_OBJECT : 0));
            $header = 'application/json';
        }
{{end}}
        if ($type == 'form') {
            // Encoding body into form-urlencoded format
            return $request->addPostFields($body);
        } else {
            // Raw body
            return $request->setBody($body, $header);
        }
    }

}
