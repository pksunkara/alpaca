var util = require('util');

/**
 * ClientError is used when the api returns an error
 */
function ClientError(message, code) {
  var err = Error.call(this, message);
  err.code = code;

  return err
}

util.inherits(ClientError, Error);

module.exports = ClientError;
