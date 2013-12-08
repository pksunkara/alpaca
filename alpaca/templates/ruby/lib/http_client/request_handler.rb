module {{.Pkg.name}}

  module HttpClient

    # RequestHandler takes care of encoding the request body into format given by options
    class RequestHandler

      def self.set_body(body, headers, options)
        flag = false
        type = options.has_key?(:request_type) ? options[:request_type] : "{{or .Api.request.formats.default "raw"}}"
{{if .Api.request.formats.json}}
        # Encoding request body into JSON format
        if type == 'json'
          body = body.to_json
          headers['Content-Type'] = 'application/json'
        end
{{end}}
        # Raw body
        if type == 'raw'
          headers.delete 'Content-Type'
        end

        return [body, headers]
      end

    end

  end

end
