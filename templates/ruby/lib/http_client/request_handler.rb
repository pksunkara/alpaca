module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # RequestHandler takes care of encoding the request body into format given by options
    class RequestHandler

      def self.set_body(options)
        type = options.has_key?(:request_type) ? options[:request_type] : "{{or .Api.request.formats.default "raw"}}"
{{if .Api.request.formats.json}}
        # Encoding request body into JSON format
        if type == "json"
          options[:body] = options[:body].to_json
          options[:headers]["content-type"] = "application/json"
        end
{{end}}
        # Encoding body into form-urlencoded format
        if type == "form"
          options[:body] = Faraday::Utils::ParamsHash[options[:body]].to_query
          options[:headers]["content-type"] = "application/x-www-form-urlencoded"
        end

        # Raw body
        if type == "raw"
          options[:body] = options[:body].is_a?(Hash) ? "" : options[:body]
          options[:headers].delete "content-type"
        end

        return options
      end

    end

  end

end
