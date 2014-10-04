var request = module.exports;

/**
 * This module takes care of encoding the request body into format given by options
 */
request.setBody = function(reqobj, body, options) {
  var type = (options['request_type'] ? options['request_type'] : '{{or .Api.Request.Formats.Default "raw"}}');

{{if .Api.Request.Formats.Json}}
  // Encoding body into JSON format
  if (type == 'json') {
    reqobj['json'] = body;
  }
{{end}}{{if .Api.Request.Formats.Form}}
  // Encoding body into form-urlencoded format
  if (type == 'form') {
    reqobj['form'] = body;
  }
{{end}}
  // Raw body
  if (type == 'raw') {
    reqobj['body'] = body;

    if (typeof reqobj['body'] == 'object') {
      delete reqobj['body'];
    }
  }

  return reqobj;
};
