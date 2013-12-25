import json

# RequestHandler takes care of encoding the request body into format given by options
class RequestHandler():

	@staticmethod
	def set_body(request):
		typ = request['request_type'] if 'request_type' in request else '{{or .Api.request.formats.default "raw"}}'
{{if .Api.request.formats.json}}
		# Encoding request body into JSON format
		if typ == 'json':
			request['data'] = json.dumps(request['data'])
			request['headers']['content-type'] = 'application/json'
{{end}}
		# Encoding body into form-urlencoded format
		if typ == 'form':
			request['headers']['content-type'] = 'application/x-www-form-urlencoded'

		if typ == 'raw':
			if 'content-type' in request['headers']:
				del request['headers']['content-type']

		if 'request_type' in request:
			del request['request_type']

		return request
