"""ClientException is used when the API returns an error"""
class ClientError(Exception):

    def __init__(self, message, code):
        super(ClientError, self).__init__()
        self.message = message
        self.code = code
