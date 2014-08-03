class ResponseHandler(object):

    """ResponseHandler takes care of decoding the response body into suitable type"""

    @staticmethod
    def get_body(response):
        typ = response.headers.get('content-type')
        body = response.text
{{if .Api.Response.Formats.Json}}
        # Response body is in JSON
        if typ.find('json') != -1:
            body = response.json()
{{end}}
        return body
