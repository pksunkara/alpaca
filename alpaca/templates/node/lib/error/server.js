var util = require("util");

function ServerError(message, code) {
  var err = Error.call(this, message);
  err.code = code;

  return err
}

util.inherits(ServerError, Error);

module.exports = ServerError;
