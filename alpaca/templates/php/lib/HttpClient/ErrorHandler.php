<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Common\Event;
use Guzzle\Http\Message\Response;

use {{.Pkg.name}}\HttpClient\ResponseHandler;
use {{.Pkg.name}}\Exception\ClientException;

/**
 * ErrorHanlder takes care of selecting the error message from response body
 */
class ErrorHandler
{
    public function onRequestError(Event $event)
    {
        $request = $event['request'];
        $response = $request->getResponse();

        $content = ResponseHandler::getBody($response);

        if ($response->isClientError() || $response->isServerError()) {
            $error = null;

            // If HTML, whole body is taken
            if (gettype($content) == "string") {
                $error = new ClientException($content, $response->getStatusCode());
            }
{{if .Api.response.formats.json}}
            // If JSON, a particular field is taken and used
            if ($response->isContentType('json') && is_array($content) && isset($content['{{.Api.error.message}}'])) {
                $error = new ClientException($content['{{.Api.error.message}}'], $response->getStatusCode());
            } else {
                $error = new ClientException("Unable to select error message from json returned by request responsible for error", $response->getStatusCode());
            }
{{end}}
            if (empty($error)) {
                $error = new \RuntimeException("Unable to understand the content type of response returned by request responsible for error", $response->getStatusCode());
            }
        }
    }
}
