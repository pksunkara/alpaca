<?php

require 'vendor/autoload.php';

set_include_path(get_include_path() . PATH_SEPARATOR . "tests/php/lib");

spl_autoload_register(function ($class) {
    $file = preg_replace('#\\\|_(?!.+\\\)#','/', $class) . '.php';

    if (stream_resolve_include_path($file)) {
        require $file;
    }
});

// Client Options

$client = new Testing\Client();

$client->clientOptions()->basic();

$tmp = new Testing\Client(array(), array(
  'base' => 'http://localhost:3001/useless',
  'api_version' => 'v2',
  'user_agent' => 'testing (user agent)',
  'headers' => array( 'custom-header' => 'custom' )
));

$tmp->clientOptions()->basic();

// Request Options

$client->requestOptions()->basic(array(
  'base' => 'http://localhost:3001/useless',
  'api_version' => 'v2',
  'headers' => array(
    'custom-header' => 'custom',
    'user-Agent' => 'testing again'
  )
));

$client->requestOptions()->suffix(array( 'response_type' => 'png' ));

// GET Request

$client->get()->api('foo', 'bar');
$client->get()->options(array( 'query' => array( 'foo' => 'bar' )));

// Responses

$response = $client->response()->basic();

$client->test()->equal(array( 'query' => array(
  'expected' => $response->code,
  'actual' => 206,
  'name' => 'The status code is correctly propogated'
)));

$client->test()->equal(array( 'query' => array(
  'expected' => $response->body,
  'actual' => 'is a response',
  'name' => 'The response body is correctly propogated'
)));

$client->test()->equal(array( 'query' => array(
  'expected' => $client->response()->header()->headers['awesome'],
  'actual' => 'wow nice',
  'name' => 'The response headers are correctly propogated'
)));

$client->test()->equal(array( 'query' => array(
  'expected' => $client->response()->html()->body,
  'actual' => 'checking html',
  'name' => 'The response body in HTML format is correctly parsed'
)));

$client->test()->equal(array( 'query' => array(
  'expected' => $client->response()->json()->body['message'],
  'actual' => 'checking json',
  'name' => 'The response body in JSON format is correctly parsed'
)));

// POST Request

$client->post()->emptyRaw();
$client->post()->optionsRaw(array( 'body' => 'hello world' ));

$client->post()->emptyForm(array( 'request_type' => 'form' ));
$client->post()->apiForm('foo', 'bar', array( 'request_type' => 'form' ));
$client->post()->optionsForm(array(
  'request_type' => 'form',
  'body' => array( 'foo' => 'bar' )
));
$client->post()->arrayForm(array(
  'request_type' => 'form',
  'body' => array( 'foo' => array('bar', 'baz') )
));

$client->post()->emptyJson(array( 'request_type' => 'json' ));
$client->post()->apiJson('foo', 'bar', array( 'request_type' => 'json' ));
$client->post()->optionsJson(array(
  'request_type' => 'json',
  'body' => array( 'foo' => array('bar', 'baz') )
));

// HTTP Methods

$client->methods()->patch();
$client->methods()->put();
$client->methods()->delete();

// Api Paths

$client->paths('lol')->basic();
$client->paths('lol')->noArg();

// Authorization

$tmp = new Testing\Client(array( 'username' => 'nine', 'password' => 'time' ));
$tmp->auth()->basic();

$tmp = new Testing\Client(array( 'http_header' => 'passwordtoken' ));
$tmp->auth()->header();

$tmp = new Testing\Client(array( 'client_id' => 'fine', 'client_secret' => 'line' ));
$tmp->auth()->oauthSecret();

$tmp = new Testing\Client('accesstoken');
$tmp->auth()->oauthToken();

// End Tests

$client->test()->end();
