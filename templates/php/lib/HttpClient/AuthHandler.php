<?php

namespace {{call .Fnc.camelize .Pkg.Name}}\HttpClient;

use Guzzle\Common\Event;

/**
 * AuthHandler takes care of devising the auth type and using it
 */
class AuthHandler
{
    private $auth;
{{if .Api.authorization.basic}}
    const HTTP_PASSWORD = 0;
{{end}}{{if .Api.authorization.header}}
    const HTTP_HEADER = 1;
{{end}}{{if .Api.authorization.oauth}}
    const URL_SECRET = 2;
    const URL_TOKEN = 3;
{{end}}
    public function __construct(array $auth = array())
    {
        $this->auth = $auth;
    }

    /**
     * Calculating the Authentication Type
     */
    public function getAuthType()
    {
{{if .Api.authorization.basic}}
        if (isset($this->auth['username']) && isset($this->auth['password'])) {
            return self::HTTP_PASSWORD;
        }
{{end}}{{if .Api.authorization.header}}
        if (isset($this->auth['http_header'])) {
            return self::HTTP_HEADER;
        }
{{end}}{{if .Api.authorization.oauth}}
        if (isset($this->auth['client_id']) && isset($this->auth['client_secret'])) {
            return self::URL_SECRET;
        }

        if (isset($this->auth['access_token'])) {
            return self::URL_TOKEN;
        }
{{end}}
        return -1;
    }

    public function onRequestBeforeSend(Event $event)
    {
        if (empty($this->auth)) {
            return;
        }

        $auth = $this->getAuthType();
        $flag = false;
{{if .Api.authorization.basic}}
        if ($auth == self::HTTP_PASSWORD) {
            $this->httpPassword($event);
            $flag = true;
        }
{{end}}{{if .Api.authorization.header}}
        if ($auth == self::HTTP_HEADER) {
            $this->httpHeader($event);
            $flag = true;
        }
{{end}}{{if .Api.authorization.oauth}}
        if ($auth == self::URL_SECRET) {
            $this->urlSecret($event);
            $flag = true;
        }

        if ($auth == self::URL_TOKEN) {
            $this->urlToken($event);
            $flag = true;
        }
{{end}}
        if (!$flag) {
            throw new \ErrorException('Unable to calculate authorization method. Please check.');
        }
    }
{{if .Api.authorization.basic}}
    /**
     * Basic Authorization with username and password
     */
    public function httpPassword(Event $event)
    {
        $event['request']->setHeader('Authorization', sprintf('Basic %s', base64_encode($this->auth['username'] . ':' . $this->auth['password'])));
    }
{{end}}{{if .Api.authorization.header}}
    /**
     * Authorization with HTTP header
     */
    public function httpHeader(Event $event)
    {
        $event['request']->setHeader('Authorization', sprintf('token %s', $this->auth['http_header']));
    }
{{end}}{{if .Api.authorization.oauth}}
    /**
     * OAUTH2 Authorization with client secret
     */
    public function urlSecret(Event $event)
    {
        $query = $event['request']->getQuery();

        $query->set('client_id', $this->auth['client_id']);
        $query->set('client_secret', $this->auth['client_secret']);
    }

    /**
     * OAUTH2 Authorization with access token
     */
    public function urlToken(Event $event)
    {
        $query = $event['request']->getQuery();

        $query->set('access_token', $this->auth['access_token']);
    }
{{end}}
}
