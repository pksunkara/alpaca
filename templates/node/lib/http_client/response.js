/**
 * Response object contains the response returned by the client
 */
var Response = function(body, code, headers) {
  this.body = body;
  this.code = code;
  this.headers = headers;
};

// Export module
module.exports = Response;
