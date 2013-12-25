// ----------------------------------------------
// npm install request 2.29.x
// npm install async
// ----------------------------------------------

var testing = require('./tests/node/lib')
  , async = require('async')
  , tasks = []
  , client;

// Client Options

client = testing.client()

tasks.push(function(callback) {
  client.clientOptions().basic(function() {
    callback();
  });
});

tasks.push(function(callback) {
  testing.client({}, {
    base: 'http://localhost:3001/useless',
    api_version: 'v2',
    user_agent: 'testing (user agent)',
    headers: { 'custom-header': 'custom' }
  }).clientOptions().basic(function() {
    callback();
  });
});

// Request Options

tasks.push(function(callback) {
  client.requestOptions().basic({
    base: 'http://localhost:3001/useless',
    api_version: 'v2',
    headers: {
      'custom-header': 'custom',
      'user-Agent': 'testing again'
    }
  }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.requestOptions().suffix({ response_type: 'png' }, function() {
    callback();
  });
});

// GET Request

tasks.push(function(callback) {
  client.get().api('foo', 'bar', function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.get().options({ query: { foo: 'bar' } }, function() {
    callback();
  });
});

// Responses

tasks.push(function(callback) {
  client.response().basic(function(err, response) {
    client.test().equal({
      query: {
        expected: response.code,
        actual: 206,
        name: 'The status code is correctly propogated'
      }
    }, function() {
      client.test().equal({
        query: {
          expected: response.body,
          actual: 'is a response',
          name: 'The response body is correctly propogated'
        }
      }, function() {
        callback();
      });
    });
  });
});

tasks.push(function(callback) {
  client.response().header(function(err, response) {
    client.test().equal({
      query: {
        expected: response.headers['awesome'],
        actual: 'wow nice',
        name: 'The response headers are correctly propogated'
      }
    }, function() {
      callback();
    });
  });
});

tasks.push(function(callback) {
  client.response().html(function(err, response) {
    client.test().equal({
      query: {
        expected: response.body,
        actual: 'checking html',
        name: 'The response body in HTML format is correctly parsed'
      }
    }, function() {
      callback();
    });
  });
});

tasks.push(function(callback) {
  client.response().json(function(err, response) {
    client.test().equal({
      query: {
        expected: response.body.message,
        actual: 'checking json',
        name: 'The response body in JSON format is correctly parsed'
      }
    }, function() {
      callback();
    });
  });
});

// POST Request

tasks.push(function(callback) {
  client.post().emptyRaw(function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().optionsRaw({ body: 'hello world' }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().emptyForm({ request_type: 'form' }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().apiForm('foo', 'bar', { request_type: 'form' }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().optionsForm({
    body: { foo: ['bar', 'baz'] },
    request_type: 'form'
  }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().emptyJson({ request_type: 'json' }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().apiJson('foo', 'bar', { request_type: 'json' }, function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.post().optionsJson({
    body: { foo: ['bar', 'baz'] },
    request_type: 'json'
  }, function() {
    callback();
  });
});

// HTTP Methods

tasks.push(function(callback) {
  client.methods().patch(function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.methods().put(function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.methods().delete(function() {
    callback();
  });
});

// Api Paths

tasks.push(function(callback) {
  client.paths('lol').basic(function() {
    callback();
  });
});

tasks.push(function(callback) {
  client.paths('lol').noArg(function() {
    callback();
  });
});

// Authorization

tasks.push(function(callback) {
  testing.client({ username: 'nine', password: 'time' }).auth().basic(function() {
    callback();
  });
});

tasks.push(function(callback) {
  testing.client({ http_header: 'passwordtoken' }).auth().header(function() {
    callback();
  });
});

tasks.push(function(callback) {
  testing.client({ client_id: 'fine', client_secret: 'line' }).auth().oauthSecret(function() {
    callback();
  });
});

tasks.push(function(callback) {
  testing.client('accesstoken').auth().oauthToken(function() {
    callback();
  });
});

// End tests

tasks.push(function(callback) {
  client.test().end(function() {
    callback();
  });
});

// Run all the tests
async.series(tasks)
