module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # ResponseHandler takes care of decoding the response body into suitable type
    class ResponseHandler

      def self.get_body(response)
        type = response.response["content-type"]
        body = response.body
{{if .Api.Response.Formats.Json}}
        # Response body is in JSON
        if type && type.include?("json")
          body = JSON.parse body
        end
{{end}}
        return body
      end

    end

  end

end
