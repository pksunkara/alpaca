# ResponseHandler takes care of decoding the response body into suitable type
class ResponseHandler():

	@staticmethod
	def get_body(response):
		return response.text
