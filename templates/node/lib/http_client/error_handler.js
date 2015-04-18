var errors = require('../error');

/**
 * ErrorHanlder takes care of selecting the error message from response body
 */
module.exports = function(response, body, callback) {
  var code = response.statusCode
    , message = ''
    , type = response.headers['content-type'];

  if (Math.floor(code/100) === 5) {
    return callback(new errors.ClientError('Error ' + code, code));
  }

  if (Math.floor(code/100) === 4) {
    // If HTML, whole body is taken
    if (typeof body === 'string') {
      message = body;
    }
{{if .Api.Response.Formats.Json}}
    // If JSON, a particular field is taken and used
    if (type.indexOf('json') !== -1 && typeof body === 'object') {
      if (body['{{.Api.Error.Message}}']) {
        message = body['{{.Api.Error.Message}}'];
      } else {
        message = 'Unable to select error message from json returned by request responsible for error';
      }
    }
{{end}}
    if (message === '') {
      message = 'Unable to understand the content type of response returned by request responsible for error';
    }

    return callback(new errors.ClientError(message, code));
  }

  return callback(null, response, body);
};
