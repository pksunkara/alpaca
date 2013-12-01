var request = module.exports;

/**
 * This module takes care of encoding the request body into format given by options
 */
request.setBody = function(reqobj, body, options) {
  var flag = false, type = (options['request_type'] ? options['request_type'] : "{{.Api.request.formats.default}}");

{{if .Api.request.formats.json}}
  // Encoding request body into JSON format
  if (type == "json") {
    flag = true;
    reqobj['json'] = body;
  }
{{end}}
  if (type == "form") {
    flag = true;
    reqobj['form'] = body;
  }

  if (!flag) {
    reqobj['body'] = body;
  }

  return reqobj;
};
