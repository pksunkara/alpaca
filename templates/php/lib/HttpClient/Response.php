<?php

namespace {{.Pkg.name}}\HttpClient;

/*
 * Response object contains the response returned by the client
 */
class Response
{
    function __construct($body, $code, $headers) {
        $this->body = $body;
        $this->code = $code;
        $this->headers = $headers;
    }
}
