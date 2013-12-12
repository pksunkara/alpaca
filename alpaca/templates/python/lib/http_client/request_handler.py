# RequestHandler takes care of encoding the request body into format given by options
class RequestHandler():

	@staticmethod
	def set_body(request):
		body = request['data']

		return request
