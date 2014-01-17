module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # ResponseHandler takes care of decoding the response body into suitable type
    class ResponseHandler

      def self.get_body(response)
        type = response.headers["content-type"]
        body = response.body
{{if .Api.response.formats.json}}
        # Response body is in JSON
        if type.include? "json"
          body = JSON.parse body
        end
{{end}}
        body
      end

    end

  end

end
