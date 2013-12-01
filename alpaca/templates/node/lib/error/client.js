var util = require("util");

function ClientError(message, code) {
  var err = Error.call(this, message);
  err.code = code;

  return err
}

util.inherits(ClientError, Error);

module.exports = ClientError;
