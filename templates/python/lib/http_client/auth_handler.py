class AuthHandler(object):
    """AuthHandler takes care of devising the auth type and using it"""
{{if .Api.authorization.basic}}
    HTTP_PASSWORD = 0
{{end}}{{if .Api.authorization.header}}
    HTTP_HEADER = 1
{{end}}{{if .Api.authorization.oauth}}
    URL_SECRET = 2
    URL_TOKEN = 3
{{end}}
    def __init__(self, auth):
        self.auth = auth

    def get_auth_type(self):
        """Calculating the Authentication Type"""
{{if .Api.authorization.basic}}
        if 'username' in self.auth and 'password' in self.auth:
            return self.HTTP_PASSWORD
{{end}}{{if .Api.authorization.header}}
        if 'http_header' in self.auth:
            return self.HTTP_HEADER
{{end}}{{if .Api.authorization.oauth}}
        if 'client_id' in self.auth and 'client_secret' in self.auth:
            return self.URL_SECRET

        if 'access_token' in self.auth:
            return self.URL_TOKEN
{{end}}
        return -1

    def set(self, request):
        if len(self.auth.keys()) == 0:
            return request

        auth = self.get_auth_type()
        flag = False
{{if .Api.authorization.basic}}
        if auth == self.HTTP_PASSWORD:
            request = self.http_password(request)
            flag = True
{{end}}{{if .Api.authorization.header}}
        if auth == self.HTTP_HEADER:
            request = self.http_header(request)
            flag = True
{{end}}{{if .Api.authorization.oauth}}
        if auth == self.URL_SECRET:
            request = self.url_secret(request)
            flag = True

        if auth == self.URL_TOKEN:
            request = self.url_token(request)
            flag = True
{{end}}
        if not flag:
            raise StandardError("Unable to calculate authorization method. Please check")

        return request
{{if .Api.authorization.basic}}
    def http_password(self, request):
        """Basic Authorization with username and password"""
        request['auth'] = (self.auth['username'], self.auth['password'])
        return request
{{end}}{{if .Api.authorization.header}}
    def http_header(self, request):
        """Authorization with HTTP header"""
        request['headers']['Authorization'] = '{{or .Api.authorization.header_prefix "token"}} ' + self.auth['http_header']
        return request
{{end}}{{if .Api.authorization.oauth}}
    def url_secret(self, request):
        """OAUTH2 Authorization with client secret"""
        request['params']['client_id'] = self.auth['client_id']
        request['params']['client_secret'] = self.auth['client_secret']
        return request

    def url_token(self, request):
        """OAUTH2 Authorization with access token"""
        request['params']['access_token'] = self.auth['access_token']
        return request
{{end}}
