var url = require('url')
  , request = require('request');

var client = module.exports;

client.AuthHandler = require('./auth_handler');
client.ErrorHandler = require('./error_handler');

client.RequestHandler = require('./request_handler');
client.ResponseHandler = require('./response_handler');

client.Response = require('./response.js');

/**
 * Main HttpClient which is used by Api classes
 */
client.HttpClient = function (auth, options) {
  if (!options) {
    options = {};
  }

  if (!auth) {
    auth = {};
  }

{{if .Api.authorization.oauth}}
  if (typeof auth == 'string') {
    auth = { 'access_token': auth };
  }
{{else}}{{if .Api.authorization.header}}
  if (typeof auth == 'string') {
    auth = { 'http_header': auth };
  }
{{end}}{{end}}
  this.options = {
    'base': '{{.Api.base}}',{{with .Api.version}}
    'api_version': '{{.}}',{{end}}
    'user_agent': 'alpaca/{{.Api.alpaca_version}} (https://github.com/pksunkara/alpaca)'
  };

  for (var key in options) {
    this.options[key] = options[key];
  }

  this.base = this.options['base'];

  this.headers = {
    'user-agent': this.options['user_agent']
  };

  if (this.options['headers']) {
    for (var key in this.options['headers']) {
      this.headers[key.toLowerCase()] = this.options['headers'][key];
    }

    delete this.options['headers'];
  }

  this.auth = new client.AuthHandler(auth);

  return this;
}

client.HttpClient.prototype.get = function (path, params, options, callback) {
  options['query'] = params;

  this.request(path, {}, 'GET', options, callback);
};

client.HttpClient.prototype.post = function (path, body, options, callback) {
  this.request(path, body, 'POST', options, callback);
};

client.HttpClient.prototype.patch = function (path, body, options, callback) {
  this.request(path, body, 'PATCH', options, callback);
};

client.HttpClient.prototype.delete = function (path, body, options, callback) {
  this.request(path, body, 'DELETE', options, callback);
};

client.HttpClient.prototype.put = function (path, body, options, callback) {
  this.request(path, body, 'PUT', options, callback);
};

/**
 * Intermediate function which does three main things
 *
 * - Transforms the body of request into correct format
 * - Creates the requests with give parameters
 * - Returns response body after parsing it into correct format
 */
client.HttpClient.prototype.request = function (path, body, method, options, callback) {
  var headers = {}, self = this;

  for (var key in this.options) {
    if (!options[key]) {
      options[key] = this.options[key];
    }
  }

  if (options['headers']) {
    headers = options['headers'];
    delete options['headers'];
  }

  for (var key in headers) {
    lowerKey = key.toLowerCase();

    if (key != lowerKey) {
      headers[lowerKey] = headers[key];
      delete headers[key];
    }
  }

  for (var key in this.headers) {
    if (!headers[key]) {
      headers[key] = this.headers[key];
    }
  }

  var reqobj = {
    'url': path,
    'qs': options['query'] || {},
    'method': method,
    'headers': headers
  };

  delete options['query'];
  delete options['body'];

  delete options['base'];
  delete options['user_agent'];
{{if .Api.no_verify_ssl}}
  reqobj['strictSSL'] = false;
{{end}}
  if (method != 'GET') {
    reqobj = this.setBody(reqobj, body, options);
  }

  reqobj = this.auth.set(reqobj);

  reqobj = this.createRequest(reqobj, options, function(err, response, body) {
    if (err) {
      return callback(err);
    }

    self.getBody(response, body, function(err, response, body) {
      if (err) {
        return callback(err);
      }

      client.ErrorHandler(response, body, function(err, response, body) {
        if (err) {
          return callback(err);
        }

        callback(null, new client.Response(body, response.statusCode, response.headers));
      });
    });
  });
};

/**
 * Creating a request with the given arguments
 *
 * If api_version is set, appends it immediately after host
 */
client.HttpClient.prototype.createRequest = function (reqobj, options, callback) {
  var version = (options['api_version'] ? '/' + options['api_version'] : '');
{{if .Api.response.suffix}}
  // Adds a suffix (ex: ".html", ".json") to url
  var suffix = (options['response_type'] ? options['response_type'] : '{{or .Api.response.formats.default "html"}}');
  reqobj['url'] = reqobj['url'] + '.' + suffix;
{{end}}
  reqobj['url'] = url.resolve(this.base, version + reqobj['url']);

  request(reqobj, callback);
};

/**
 * Get response body in correct format
 */
client.HttpClient.prototype.getBody = function (response, body, callback) {
  client.ResponseHandler.getBody(response, body, callback);
};

/**
 * Set request body in correct format
 */
client.HttpClient.prototype.setBody = function (request, body, options) {
  return client.RequestHandler.setBody(request, body, options);
};
