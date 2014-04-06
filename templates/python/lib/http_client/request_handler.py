import urllib
import json


class RequestHandler(object):

    """RequestHandler takes care of encoding the request body into format given by options."""

    @staticmethod
    def render_key(parents):
        depth, new = 0, ''

        for x in parents:
            old = '[%s]' if depth > 0 else '%s'
            new += old % x
            depth += 1

        return new

    @staticmethod
    def url_encode(data, parents=None, pairs=None):
        if pairs is None:
            pairs = {}

        if parents is None:
            parents = []

        if isinstance(data, dict):
            for key, value in data.items():
                RequestHandler.urlencode(value, parents + [key], pairs)
        elif isinstance(data, list):
            for key, value in enumerate(data):
                RequestHandler.urlencode(value, parents + [key], pairs)
        else:
            pairs[RequestHandler.renderKey(parents)] = data

        return pairs

    @staticmethod
    def set_body(request):
        typ = request['request_type'] if 'request_type' in request else '{{or .Api.request.formats.default "raw"}}'
{{if .Api.request.formats.json}}
        # Encoding request body into JSON format
        if typ == 'json':
            request['data'] = json.dumps(request['data'])
            request['headers']['content-type'] = 'application/json'
{{end}}
        # Encoding body into form-urlencoded format
        if typ == 'form':
            request['data'] = RequestHandler.urlencode(request['data'])
            request['headers']['content-type'] = 'application/x-www-form-urlencoded'

        if typ == 'raw':
            if 'content-type' in request['headers']:
                del request['headers']['content-type']

        if 'request_type' in request:
            del request['request_type']

        return request
