import requests
import copy

from auth_handler import AuthHandler
from error_handler import ErrorHandler
from request_handler import RequestHandler
from response import Response
from response_handler import ResponseHandler

# Main HttpClient which is used by Api classes
class HttpClient():

	def __init__(self, auth, options):
{{if .Api.authorization.oauth}}
		if isinstance(auth, str):
			auth = { 'access_token': auth }
{{else}}{{if .Api.authorization.header}}
		if isinstance(auth, str):
			auth = { 'http_header': auth }
{{end}}{{end}}
		self.options = {
			'base': '{{.Api.base}}',{{with .Api.version}}
			'api_version': '{{.}}',{{end}}
			'user_agent': 'alpaca/0.1.0 (https://github.com/pksunkara/alpaca)'
		}

		self.options.update(options)

		self.headers = {
			'user-agent': self.options['user_agent']
		}

		if 'headers' in self.options:
			self.headers.update(self.options['headers'])
			del self.options['headers']

		self.auth = AuthHandler(auth)

	def get(self, path, params={}, options={}):
		options.update({ 'query': params })
		return self.request(path, None, 'get', options)

	def post(self, path, body={}, options={}):
		return self.request(path, body, 'post', options)

	def patch(self, path, body={}, options={}):
		return self.request(path, body, 'patch', options)

	def delete(self, path, body={}, options={}):
		return self.request(path, body, 'delete', options)

	def put(self, path, body={}, options={}):
		return self.request(path, body, 'put', options)

	# Intermediate function which does three main things
	#
	# - Transforms the body of request into correct format
	# - Creates the requests with give parameters
	# - Returns response body after parsing it into correct format
	def request(self, path, body, method, options):
		kwargs = copy.deepcopy(self.options)
		kwargs.update(options)

		kwargs['headers'] = copy.deepcopy(self.headers)

		if 'headers' in options:
			kwargs['headers'].update(options['headers'])

		kwargs['data'] = body
		kwargs['allow_redirects'] = True

		kwargs['params'] = kwargs['query'] if 'query' in kwargs else {}

		if method != 'get':
			kwargs = self.set_body(kwargs)

		kwargs['hooks'] = dict(response=ErrorHandler.check_error)

		kwargs = self.auth.set(kwargs)

		response = self.create_request(method, path, kwargs)

		return Response(self.get_body(response), response.status_code, response.headers)

	# Creating a request with the given arguments
	#
	# If api_version is set, appends it immediately after host
	def create_request(method, path, options):
		version = '/' + options['api_version'] if 'api_version' in options else ''
{{if .Api.response.suffix}}
		# Adds a suffix (ex: ".html", ".json") to url
		suffix = options['response_type'] if 'response_type' in options else '{{or .Api.response.formats.default "html"}}'
		path = path + '.' + suffix
{{end}}
		path = options['base'] + version + path

		return requests.request(method, path, **options)

	# Get response body in correct format
	def get_body(self, response):
		return ResponseHandler.get_body(response)

	# Set request body in correct format
	def set_body(self, request):
		return RequestHandler.set_body(request)
