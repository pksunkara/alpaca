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

        $message = null;
        $code = $response->getStatusCode();

        $content = ResponseHandler::getBody($response);

        if ($response->isServerError()) {
            throw new ClientException('Error '.$code, $code);
        }

        if ($response->isClientError()) {
            // If HTML, whole body is taken
            if (gettype($content) == "string") {
                $message = $content;
            }
{{if .Api.response.formats.json}}
            // If JSON, a particular field is taken and used
            if ($response->isContentType('json') && is_array($content) && isset($content['{{.Api.error.message}}'])) {
                $message = $content['{{.Api.error.message}}'];
            } else {
                $message = "Unable to select error message from json returned by request responsible for error";
            }
{{end}}
            if (empty($message)) {
                $message = "Unable to understand the content type of response returned by request responsible for error";
            }

            throw new ClientException($message, $code);
        }
    }
}
