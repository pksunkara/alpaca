<?php

namespace {{call .Fnc.camelize .Pkg.Name}}\HttpClient;

use Guzzle\Http\Client as GuzzleClient;
use Guzzle\Http\ClientInterface;
use Guzzle\Http\Message\RequestInterface;

use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\AuthHandler;
use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\ErrorHandler;
use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\RequestHandler;
use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\Response;
use {{call .Fnc.camelize .Pkg.Name}}\HttpClient\ResponseHandler;

/**
 * Main HttpClient which is used by Api classes
 */
class HttpClient
{
    protected $options = array(
        'base'    => '{{.Api.base}}',{{with .Api.version}}
        'api_version' => '{{.}}',{{end}}
        'user_agent' => 'alpaca/{{.Api.alpaca_version}} (https://github.com/pksunkara/alpaca)'
    );

    protected $headers = array();

    public function __construct($auth = array(), array $options = array())
    {
{{if .Api.authorization.oauth}}
        if (gettype($auth) == 'string') {
            $auth = array('access_token' => $auth);
        }
{{else}}{{if .Api.authorization.header}}
        if (gettype($auth) == 'string') {
            $auth = array('http_header' => $auth);
        }
{{end}}{{end}}
        $this->options = array_merge($this->options, $options);

        $this->headers = array(
            'user-agent' => $this->options['user_agent'],
        );

        if (isset($this->options['headers'])) {
            $this->headers = array_merge($this->headers, array_change_key_case($this->options['headers']));
            unset($this->options['headers']);
        }

        $client = new GuzzleClient($this->options['base'], $this->options);
        $this->client = $client;

        $listener = array(new ErrorHandler(), 'onRequestError');
        $this->client->getEventDispatcher()->addListener('request.error', $listener);

        if (!empty($auth)) {
            $listener = array(new AuthHandler($auth), 'onRequestBeforeSend');
            $this->client->getEventDispatcher()->addListener('request.before_send', $listener);
        }
    }

    public function get($path, array $params = array(), array $options = array())
    {
        return $this->request($path, null, 'GET', array_merge($options, array('query' => $params)));
    }

    public function post($path, $body, array $options = array())
    {
        return $this->request($path, $body, 'POST', $options);
    }

    public function patch($path, $body, array $options = array())
    {
        return $this->request($path, $body, 'PATCH', $options);
    }

    public function delete($path, $body, array $options = array())
    {
        return $this->request($path, $body, 'DELETE', $options);
    }

    public function put($path, $body, array $options = array())
    {
        return $this->request($path, $body, 'PUT', $options);
    }

    /**
     * Intermediate function which does three main things
     *
     * - Transforms the body of request into correct format
     * - Creates the requests with give parameters
     * - Returns response body after parsing it into correct format
     */
    public function request($path, $body = null, $httpMethod = 'GET', array $options = array())
    {
        $headers = array();

        $options = array_merge($this->options, $options);

        if (isset($options['headers'])) {
            $headers = $options['headers'];
            unset($options['headers']);
        }

        $headers = array_merge($this->headers, array_change_key_case($headers));

        unset($options['body']);

        unset($options['base']);
        unset($options['user_agent']);
{{if .Api.no_verify_ssl}}
        $options['verify'] = false;
{{end}}
        $request = $this->createRequest($httpMethod, $path, null, $headers, $options);

        if ($httpMethod != 'GET') {
            $request = $this->setBody($request, $body, $options);
        }

        try {
            $response = $this->client->send($request);
        } catch (\LogicException $e) {
            throw new \ErrorException($e->getMessage());
        } catch (\RuntimeException $e) {
            throw new \RuntimeException($e->getMessage());
        }

        return new Response($this->getBody($response), $response->getStatusCode(), $response->getHeaders());
    }

    /**
     * Creating a request with the given arguments
     *
     * If api_version is set, appends it immediately after host
     */
    public function createRequest($httpMethod, $path, $body = null, array $headers = array(), array $options = array())
    {
        $version = (isset($options['api_version']) ? "/".$options['api_version'] : "");
{{if .Api.response.suffix}}
        // Adds a suffix (ex: ".html", ".json") to url
        $suffix = (isset($options['response_type']) ? $options['response_type'] : "{{or .Api.response.formats.default "html"}}");
        $path = $path.".".$suffix;
{{end}}
        $path    = $version.$path;

        return $this->client->createRequest($httpMethod, $path, $headers, $body, $options);
    }

    /**
     * Get response body in correct format
     */
    public function getBody($response)
    {
        return ResponseHandler::getBody($response);
    }

    /**
     * Set request body in correct format
     */
    public function setBody(RequestInterface $request, $body, $options)
    {
        return RequestHandler::setBody($request, $body, $options);
    }
}
