var errors = require('../error');

/**
 * ErrorHanlder takes care of selecting the error message from response body
 */
module.exports = function(response, body, callback) {
  var code = response.statusCode, message = '';

  if (Math.floor(code/100) == 5) {
    return callback(new errors.ClientError('Error ' + code, code)));
  }

  // If HTML, whole body is taken
  if (typeof body == 'string') {
    message = body;
  }
{{if .Api.response.formats.json}}
  // If JSON, a particular field is taken and used
  if (body['{{.Api.error.message}}']) {
    message = body['{{.Api.error.message}}'];
  } else {
    message = 'Unable to select error message from json returned by request responsible for error';
  }
{{end}}
  if (message == '') {
    message = 'Unable to understand the content type of response returned by request responsible for error';
  }

  if (Math.floor(code/100) == 4) {
    return callback(new errors.ClientError(message, code));
  }

  return callback(null, response, body);
};
