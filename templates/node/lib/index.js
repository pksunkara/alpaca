var Client = require('./{{call .Fnc.camelizeDownFirst .Pkg.Name}}/client');

// Export module
var {{call .Fnc.camelizeDownFirst .Pkg.Name}} = module.exports;

/**
 * This file contains the global namespace for the module
 */
{{call .Fnc.camelizeDownFirst .Pkg.Name}}.client = function(auth, options) {
  return new Client(auth, options);
};
