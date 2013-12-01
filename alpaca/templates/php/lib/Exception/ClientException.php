<?php

namespace {{.Pkg.name}}\Exception;

/**
 * ClientException is used when the api returns an error
 */
class ClientException extends \ErrorException implements ExceptionInterface
{

}
