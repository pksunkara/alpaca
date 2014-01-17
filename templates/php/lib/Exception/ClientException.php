<?php

namespace {{call .Fnc.camelize .Pkg.Name}}\Exception;

/**
 * ClientException is used when the api returns an error
 */
class ClientException extends \ErrorException implements ExceptionInterface
{

    public $code = null;

    public function __construct($message, $code) {
        $this->code = $code;
        parent::__construct($message);
    }

}
