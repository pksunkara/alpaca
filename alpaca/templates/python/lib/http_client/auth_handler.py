# AuthHandler takes care of devising the auth type and using it
class AuthHandler():

	HTTP_PASSWORD = 0
	HTTP_TOKEN = 1
{{if .Api.authorization.oauth}}
	URL_SECRET = 2
	URL_TOKEN = 3
{{end}}
	def __init__(self, auth):
		self.auth = auth

	# Calculating the Authentication Type
	def get_auth_type(self):
		if 'username' in self.auth and 'password' in self.auth:
			return self.HTTP_PASSWORD
		elif 'http_token' in self.auth:
			return self.HTTP_TOKEN{{if .Api.authorization.oauth}}
		elif 'client_id' in self.auth and 'client_secret' in self.auth:
			return self.URL_SECRET
		elif 'access_token' in self.auth:
			return self.URL_TOKEN{{end}}
		else:
			return -1

	def set(self, request):
		if len(self.auth.keys()) == 0:
			return request

		auth = self.get_auth_type()

		if auth == HTTP_PASSWORD:
			request = http_password(request)
		elif auth == HTTP_TOKEN:
			request = http_token(request){{if .Api.authorization.oauth}}
		elif auth == URL_SECRET:
			request = url_secret(request)
		elif auth == URL_TOKEN:
			request = url_token(request){{end}}
		else:
			raise StandardError("Unable to calculate authorization method. Please check")

		return request

	# Basic Authorization with username and password
	def http_password(self, request):
		request['auth'] = (self.auth['username'], self.auth['password'])
		return request

	# Authorization with HTTP token
	def http_token(self, request):
		request['headers']['Authorization'] = 'token ' + self.auth['http_token']
		return request
{{if .Api.authorization.oauth}}
	# OAUTH2 Authorization with client secret
	def url_secret(self, request):
		request['params']['client_id'] = self.auth['client_id']
		request['params']['client_secret'] = self.auth['client_secret']
		return request

	# OAUTH2 Authorization with access token
	def url_token(self, request):
		request['params']['access_token'] = self.auth['access_token']
		return request
{{end}}
