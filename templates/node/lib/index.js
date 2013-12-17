var Client = require('./{{call .Fnc.camelizeDownFirst .Pkg.name}}/client');

// Export module
var {{call .Fnc.camelizeDownFirst .Pkg.name}} = module.exports;

/**
 * This file contains the global namespace for the module
 */
{{call .Fnc.camelizeDownFirst .Pkg.name}}.client = function(auth, options) {
  return new Client(auth, options);
};
