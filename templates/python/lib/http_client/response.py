# Response object contains the response returned by the client
class Response():
	def __init__(self, body, code, headers):
		self.body = body
		self.code = code
		self.headers = headers
