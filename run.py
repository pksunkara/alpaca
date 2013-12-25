import os, sys
sys.path.append(os.path.abspath('./tests/python'))

import testing

# Client Options

client = testing.Client()

client.client_options().basic()

testing.Client({}, {
  'base': 'http://localhost:3001/useless',
  'api_version': 'v2',
  'user_agent': 'testing (user agent)',
  'headers': {
    'custom-header': 'custom'
  }
}).client_options().basic()

# Request Options

client.request_options().basic({
  'base': 'http://localhost:3001/useless',
  'api_version': 'v2',
  'headers': {
    'custom-header': 'custom',
    'user-Agent': 'testing again'
  }
})

client.request_options().suffix({ 'response_type': 'png' })

# GET Request

client.get().api('foo', 'bar')
client.get().options({ 'query': { 'foo': 'bar' }})

# Responses

response = client.response().basic()

client.test().equal({ 'query': {
  'expected': response.code,
  'actual': 206,
  'name': 'The status code is correctly propogated'
}})

client.test().equal({ 'query': {
  'expected': response.body,
  'actual': 'is a response',
  'name': 'The response body is correctly propogated'
}})

client.test().equal({ 'query': {
  'expected': client.response().header().headers['awesome'],
  'actual': 'wow nice',
  'name': 'The response headers are correctly propogated'
}})

client.test().equal({ 'query': {
  'expected': client.response().html().body,
  'actual': 'checking html',
  'name': 'The response body in HTML format is correctly parsed'
}})

client.test().equal({ 'query': {
  'expected': client.response().json().body['message'],
  'actual': 'checking json',
  'name': 'The response body in JSON format is correctly parsed'
}})

# POST Request

client.post().empty_raw()
client.post().options_raw({ 'body': 'hello world' })

client.post().empty_form({ 'request_type': 'form' })
client.post().api_form('foo', 'bar', { 'request_type': 'form' })
client.post().options_form({
  'request_type': 'form',
  'body': { 'foo': ['bar', 'baz'] }
})

client.post().empty_json({ 'request_type': 'json' })
client.post().api_json('foo', 'bar', { 'request_type': 'json' })
client.post().options_json({
  'request_type': 'json',
  'body': { 'foo': ['bar', 'baz'] }
})

# HTTP Methods

client.methods().patch()
client.methods().put()
client.methods().delete()

# Api Paths

client.paths('lol').basic()
client.paths('lol').no_arg()

# Authorization

testing.Client({ 'username': 'nine', 'password': 'time' }).auth().basic()
testing.Client({ 'http_header': 'passwordtoken' }).auth().header()
testing.Client({ 'client_id': 'fine', 'client_secret': 'line' }).auth().oauth_secret()
testing.Client('accesstoken').auth().oauth_token()

# End tests

client.test().end()
