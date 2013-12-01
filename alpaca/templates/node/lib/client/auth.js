/**
 * This module takes care of devising the auth type and using it
 */
var Auth = function(auth) {
  this.auth = auth;

  this.HTTP_PASSWORD = 0;
  this.HTTP_TOKEN = 1;
{{if .Api.authorization.oauth}}
  this.URL_SECRET = 2;
  this.URL_TOKEN = 3;
{{end}}
  return this;
};

/**
 * Calculating the type of authentication
 */
Auth.prototype.getAuthType = function () {
  if (this.auth['username'] && this.auth['password']) {
    return HTTP_PASSWORD;
  } else if (this.auth['http_token']) {
    return HTTP_TOKEN;
  }{{if .Api.authorization.oauth}} else if (this.auth['client_id'] && this.auth['client_secret']) {
    return URL_SECRET;
  } else if (this.auth['access_token']) {
    return URL_TOKEN;
  }{{end}}
};

/**
 * Set authentication for the request
 *
 * This should throw error because this should be caught while in development
 */
Auth.prototype.set = function (request) {
  if (Object.keys(this.auth).length == 0) {
    return request;
  }

  switch (this.getAuthType()) {
    case this.HTTP_PASSWORD:
      request = this.http_password(request);
      break;

    case this.HTTP_TOKEN:
      request = this.http_token(request);
      break;
{{if .Api.authorization.oauth}}
    case this.URL_SECRET:
      request = this.url_secret(request);
      break;

    case this.URL_TOKEN:
      request = this.url_token(request);
      break;
{{end}}
    default:
      throw new Error('Unable to calculate authorization method. Please check.');
  }

  return request;
};

/**
 * Basic Authorization with username and password
 */
Auth.prototype.http_password = function(request) {
  request['auth'] = this.auth;

  return request;
};

/**
 * Authorization with HTTP token
 */
Auth.prototype.http_token = function(request) {
  request['headers']['Authorization'] = 'token ' + this.auth['http_token'];

  return request;
};
{{if .Api.authorization.oauth}}
/**
 * OAUTH2 Authorization with client secret
 */
Auth.prototype.url_secret = function(request) {
  request['qs']['client_id'] = this.auth['client_id'];
  request['qs']['client_id'] = this.auth['client_secret'];

  return request;
};

/**
 * OAUTH2 Authorization with access token
 */
Auth.prototype.url_token = function(request) {
  request['qs']['access_token'] = this.auth['access_token'];

  return request;
};
{{end}}
// Export module
module.exports = Auth;
