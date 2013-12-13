<?php

namespace {{.Pkg.name}}\HttpClient;

use Guzzle\Common\Event;

/**
 * AuthHandler takes care of devising the auth type and using it
 */
class AuthHandler
{
    private $auth;

    const HTTP_PASSWORD = 0;
    const HTTP_TOKEN = 1;
{{if .Api.authorization.oauth}}
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
        if (isset($this->auth['username']) && isset($this->auth['password'])) {
            return self::HTTP_PASSWORD;
        } else if (isset($this->auth['http_token'])) {
            return self::HTTP_TOKEN;
        }{{if .Api.authorization.oauth}} else if (isset($this->auth['client_id']) && isset($this->auth['client_secret'])) {
            return self::URL_SECRET;
        } else if (isset($this->auth['access_token'])) {
            return self::URL_TOKEN;
        }{{end}}
    }

    public function onRequestBeforeSend(Event $event)
    {
        if (empty($this->auth)) {
            return;
        }

        switch ($this->getAuthType()) {
            case self::HTTP_PASSWORD:
                $this->httpPassword($event);
                break;

            case self::HTTP_TOKEN:
                $this->httpToken($event);
                break;
{{if .Api.authorization.oauth}}
            case self::URL_SECRET:
                $this->urlSecret($event);
                break;

            case self::URL_TOKEN:
                $this->urlToken($event);
                break;
{{end}}
            default:
                throw new \ErrorException('Unable to calculate authorization method. Please check.');
                break;
        }
    }

    /**
     * Basic Authorization with username and password
     */
    public function httpPassword(Event $event)
    {
        $event['request']->setHeader('Authorization', sprintf('Basic %s', base64_encode($this->auth['username'] . ':' . $this->auth['password'])));
    }

    /**
     * Authorization with HTTP token
     */
    public function httpToken(Event $event)
    {
        $event['request']->setHeader('Authorization', sprintf('token %s', $this->auth['http_token']));
    }
{{if .Api.authorization.oauth}}
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
