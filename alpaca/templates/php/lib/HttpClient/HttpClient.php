<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Http\Client as GuzzleClient;
use Guzzle\Http\ClientInterface;
use Guzzle\Http\Message\Request;
use Guzzle\Http\Message\Response;

use {{.Pkg.name}}\Exception\ErrorException;
use {{.Pkg.name}}\Exception\RuntimeException;
use {{.Pkg.name}}\HttpClient\AuthHandler;
use {{.Pkg.name}}\HttpClient\ErrorHandler;

class HttpClient
{
    protected $options = array(
        'base'    => '{{.Api.base}}',
{{with .Api.version}}
        'api_version' => '{{.}}',
{{end}}
        'user_agent' => 'alpaca/0.1.0 (https://github.com/pksunkara/alpaca)'
    );

    protected $headers = array();

    public function __construct($auth = array(), array $options = array())
    {
{{if .Api.authorization.oauth}}
        if (gettype($auth) == 'string') {
            $auth = array('access_token' => $auth);
        }
{{end}}
        $this->options = array_merge($this->options, $options);

        $client = new GuzzleClient($this->options['base'], $this->options);
        $this->client  = $client;

        $listener = array(new ErrorHandler(), 'onRequestError');
        $this->client->getEventDispatcher()->addListener('request.error', $listener);

        if (!empty($auth)) {
            $listener = array(new AuthHandler($auth), 'onRequestBeforeSend');
            $this->client->getEventDispatcher()->addListener('request.before_send', $listener);
        }

        $this->clearHeaders();
    }

    public function setOption($name, $value)
    {
        $this->options[$name] = $value;
    }

    public function setHeaders(array $headers)
    {
        $this->headers = array_merge($this->headers, $headers);
    }

    public function clearHeaders()
    {
        $this->headers = array(
            sprintf('User-Agent: %s', $this->options['user_agent']),
        );
    }

    public function get($path, array $parameters = array(), array $headers = array())
    {
        return $this->request($path, null, 'GET', $headers, array('query' => $parameters));
    }

    public function post($path, $body = null, array $headers = array())
    {
        return $this->request($path, $body, 'POST', $headers);
    }

    public function patch($path, $body = null, array $headers = array())
    {
        return $this->request($path, $body, 'PATCH', $headers);
    }

    public function delete($path, $body = null, array $headers = array())
    {
        return $this->request($path, $body, 'DELETE', $headers);
    }

    public function put($path, $body, array $headers = array())
    {
        return $this->request($path, $body, 'PUT', $headers);
    }

    public function request($path, $body = null, $httpMethod = 'GET', array $headers = array(), array $options = array())
    {
        $request = $this->createRequest($httpMethod, $path, $body, $headers, $options);

        try {
            $response = $this->client->send($request);
        } catch (\LogicException $e) {
            throw new ErrorException($e->getMessage());
        } catch (\RuntimeException $e) {
            throw new RuntimeException($e->getMessage());
        }

        $this->lastRequest  = $request;
        $this->lastResponse = $response;

        return $response;
    }

    public function getLastRequest()
    {
        return $this->lastRequest;
    }

    public function getLastResponse()
    {
        return $this->lastResponse;
    }

    public function createRequest($httpMethod, $path, $body = null, array $headers = array(), array $options = array())
    {
        $version = (isset($this->options['api_version']) ? "/".$this->options['api_version'] : "");
        $path    = $version.$path;
        $headers = array_merge($this->headers, $headers);

        return $this->client->createRequest($httpMethod, $path, $headers, $body, $options);
    }
}
