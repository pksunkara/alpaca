<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Common\Event;
use Guzzle\Http\Message\Response;

use {{.Pkg.name}}\HttpClient\ResponseHandler;
use {{.Pkg.name}}\Exception\ErrorException;
use {{.Pkg.name}}\Exception\RuntimeException;

/*
 * ErrorHanlder takes care of selecting the error message from response body
 */
class ErrorHandler
{
    public function onRequestError(Event $event)
    {
        $request = $event['request'];
        $response = $request->getResponse();

        if ($response->isClientError() || $response->isServerError()) {
            $content = Response::getBody($response);

            // If HTML, whole body is taken
            if (gettype($content) == "string") {
                $error = new ErrorException($content, $response->getStatusCode());
            }
{{if .Api.response.formats.json}}
            // If JSON, a particular field is taken and used
            if ($response->isContentType('json') && is_array($content) && isset($content['{{.Api.error.message}}'])) {
                $error = new ErrorException($content['{{.Api.error.message}}'], $response->getStatusCode());
            } else {
                $error = new RuntimeException(isset($content['{{.Api.error.message}}']) ? $content['{{.Api.error.message}}'] : $content, $response->getStatusCode());
            }
{{end}}
            throw $error;
        }
    }
}
