var Client = require('./{{call .Fnc.underscore .Pkg.Name}}/client');

// Export module
var {{call .Fnc.camelizeDownFirst .Pkg.Name}} = module.exports;

/**
 * This file contains the global namespace for the module
 */
{{call .Fnc.camelizeDownFirst .Pkg.Name}}.client = function({{if .Api.base_as_arg}}baseUrl, {{end}}auth, options) {
  return new Client({{if .Api.base_as_arg}}baseUrl, {{end}}auth, options);
};
