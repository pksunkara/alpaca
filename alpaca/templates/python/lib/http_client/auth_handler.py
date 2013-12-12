# AuthHandler takes care of devising the auth type and using it
class AuthHandler():

	def __init__(self, auth):
		self.auth = auth

	def set(self, request):
		return request
