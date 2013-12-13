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
    return this.HTTP_PASSWORD;
  } else if (this.auth['http_token']) {
    return this.HTTP_TOKEN;
  }{{if .Api.authorization.oauth}} else if (this.auth['client_id'] && this.auth['client_secret']) {
    return this.URL_SECRET;
  } else if (this.auth['access_token']) {
    return this.URL_TOKEN;
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
      request = this.httpPassword(request);
      break;

    case this.HTTP_TOKEN:
      request = this.httpToken(request);
      break;
{{if .Api.authorization.oauth}}
    case this.URL_SECRET:
      request = this.urlSecret(request);
      break;

    case this.URL_TOKEN:
      request = this.urlToken(request);
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
Auth.prototype.httpPassword = function(request) {
  request['auth'] = this.auth;

  return request;
};

/**
 * Authorization with HTTP token
 */
Auth.prototype.httpToken = function(request) {
  request['headers']['Authorization'] = 'token ' + this.auth['http_token'];

  return request;
};
{{if .Api.authorization.oauth}}
/**
 * OAUTH2 Authorization with client secret
 */
Auth.prototype.urlSecret = function(request) {
  request['qs']['client_id'] = this.auth['client_id'];
  request['qs']['client_id'] = this.auth['client_secret'];

  return request;
};

/**
 * OAUTH2 Authorization with access token
 */
Auth.prototype.urlToken = function(request) {
  request['qs']['access_token'] = this.auth['access_token'];

  return request;
};
{{end}}
// Export module
module.exports = Auth;
