<?php

namespace {{call .Fnc.camelize .Pkg.Name}}\HttpClient;

use Guzzle\Common\Event;
use Guzzle\Http\Message\Response;

use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\ResponseHandler;
use {{call .Fnc.camelize .Pkg.Name}}\Exception\ClientException;

/**
 * ErrorHanlder takes care of selecting the error message from response body
 */
class ErrorHandler
{
    public function onRequestError(Event $event)
    {
        $request = $event['request'];
        $response = $request->getResponse();

        $message = null;
        $code = $response->getStatusCode();

        if ($response->isServerError()) {
            throw new ClientException('Error '.$code, $code);
        }

        if ($response->isClientError()) {
            $body = ResponseHandler::getBody($response);

            // If HTML, whole body is taken
            if (gettype($body) == 'string') {
                $message = $body;
            }
{{if .Api.response.formats.json}}
            // If JSON, a particular field is taken and used
            if ($response->isContentType('json') && is_array($body)) {
                if (isset($body['{{.Api.error.message}}'])) {
                    $message = $body['{{.Api.error.message}}'];
                } else {
                    $message = 'Unable to select error message from json returned by request responsible for error';
                }
            }
{{end}}
            if (empty($message)) {
                $message = 'Unable to understand the content type of response returned by request responsible for error';
            }

            throw new ClientException($message, $code);
        }
    }
}
