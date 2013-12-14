import string

# ResponseHandler takes care of decoding the response body into suitable type
class ResponseHandler():

	@staticmethod
	def get_body(response):
		typ = response.headers.get('content-type')
		body = response.text
{{if .Api.response.formats.json}}
		# Response body is in JSON
		if string.find(typ, 'json') != -1:
			body = response.json()
{{end}}
		return body
