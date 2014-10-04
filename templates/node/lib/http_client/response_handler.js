var response = module.exports;

/**
 * This module takes care of decoding the response body into suitable type
 */
response.getBody = function(res, body, callback) {
  var type = res.headers['content-type'], error = null;
{{if .Api.Response.Formats.Json}}
  // Response body is in JSON
  if (type.indexOf('json') != -1 && typeof(body) != 'object') {
    try {
      body = JSON.parse(body || '{}');
    } catch (err) {
      error = err;
    }
  }
{{end}}
  return callback(error, res, body);
};
