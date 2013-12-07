var request = module.exports;

/**
 * This module takes care of encoding the request body into format given by options
 */
request.setBody = function(reqobj, body, options) {
  var flag = false, type = (options['request_type'] ? options['request_type'] : "{{or .Api.request.formats.default "raw"}}");

{{if .Api.request.formats.json}}
  // Encoding body into JSON format
  if (type == "json") {
    flag = true;
    reqobj['json'] = body;
    reqobj['headers']['Content-Type'] = 'application/json';
  }
{{end}}
  // Encoding body into form-urlencoded format
  if (type == "form") {
    flag = true;
    reqobj['form'] = body;
  }

  // Raw body
  if (!flag || type == "raw") {
    reqobj['body'] = body;
  }

  return reqobj;
};
