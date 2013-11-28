<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Common\Event;
use Guzzle\Http\Message\Response;

use {{.Pkg.name}}\HttpClient\Response;
use {{.Pkg.name}}\Exception\ErrorException;
use {{.Pkg.name}}\Exception\RuntimeException;

class ErrorHandler
{
    public function onRequestError(Event $event)
    {
        $request = $event['request'];
        $response = $request->getResponse();

        if ($response->isClientError() || $response->isServerError()) {
            $content = Response::getContent($response);

            if (is_array($content) && isset($content['{{.Api.error.message}}'])) {
                throw new ErrorException($content['{{.Api.error.message}}'], $response->getStatusCode());
            }

            throw new RuntimeException(isset($content['{{.Api.error.message}}']) ? $content['{{.Api.error.message}}'] : $content, $response->getStatusCode());
        }
    }
}
