/**
 * This module takes care of devising the auth type and using it
 */
var Auth = function(auth) {
  this.auth = auth;
{{if .Api.Authorization.Basic}}
  this.HTTP_PASSWORD = 0;
{{end}}{{if .Api.Authorization.Header}}
  this.HTTP_HEADER = 1;
{{end}}{{if .Api.Authorization.Oauth}}
  this.URL_SECRET = 2;
  this.URL_TOKEN = 3;
{{end}}
  return this;
};

/**
 * Calculating the type of authentication
 */
Auth.prototype.getAuthType = function () {
{{if .Api.Authorization.Basic}}
  if (this.auth.username && this.auth.password) {
    return this.HTTP_PASSWORD;
  }
{{end}}{{if .Api.Authorization.Header}}
  if (this.auth.http_header) {
    return this.HTTP_HEADER;
  }
{{end}}{{if .Api.Authorization.Oauth}}
  if (this.auth.client_id && this.auth.client_secret) {
    return this.URL_SECRET;
  }

  if (this.auth.access_token) {
    return this.URL_TOKEN;
  }
{{end}}
  return -1;
};

/**
 * Set authentication for the request
 *
 * This should throw error because this should be caught while in development
 */
Auth.prototype.set = function (request, callback) {
  if (Object.keys(this.auth).length === 0) {
    return callback({{if .Api.Authorization.NeedAuth}}new Error('Server requires authentication to proceed further. Please check'));
  {{else}}null, request);
  {{end}}}

  var auth = this.getAuthType(), flag = false;
{{if .Api.Authorization.Basic}}
  if (auth === this.HTTP_PASSWORD) {
    request = this.httpPassword(request);
    flag = true;
  }
{{end}}{{if .Api.Authorization.Header}}
  if (auth == this.HTTP_HEADER) {
    request = this.httpHeader(request);
    flag = true;
  }
{{end}}{{if .Api.Authorization.Oauth}}
  if (auth == this.URL_SECRET) {
    request = this.urlSecret(request);
    flag = true;
  }

  if (auth == this.URL_TOKEN) {
    request = this.urlToken(request);
    flag = true;
  }
{{end}}
  if (!flag) {
      return callback(new Error('Unable to calculate authorization method. Please check.'));
  }

  callback(null, request);
};
{{if .Api.Authorization.Basic}}
/**
 * Basic Authorization with username and password
 */
Auth.prototype.httpPassword = function(request) {
  request.auth = this.auth;

  return request;
};
{{end}}{{if .Api.Authorization.Header}}
/**
 * Authorization with HTTP header
 */
Auth.prototype.httpHeader = function(request) {
  request['headers']['Authorization'] = '{{or .Api.Authorization.HeaderPrefix "token"}} ' + this.auth['http_header'];

  return request;
};
{{end}}{{if .Api.Authorization.Oauth}}
/**
 * OAUTH2 Authorization with client secret
 */
Auth.prototype.urlSecret = function(request) {
  request['qs']['client_id'] = this.auth['client_id'];
  request['qs']['client_secret'] = this.auth['client_secret'];

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
