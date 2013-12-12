# ClientException is used when the api returns an error
class ClientError(StandardError):

	def __init__(self, message, code):
		super(ClientError, self).__init__()
		self.message = message
		self.code = code
