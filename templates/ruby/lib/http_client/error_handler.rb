module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # ErrorHanlder takes care of selecting the error message from response body
    class ErrorHandler < Faraday::Middleware

      def initialize(app)
        super(app)
      end

      def call(env)
        @app.call(env).on_complete do |env|
          code = env.status
          type = env.response_headers["content-type"]

          case code
          when 500...599
            raise {{call .Fnc.camelize .Pkg.Name}}::Error::ClientError.new("Error #{code}", code)
          when 400...499
            body = {{call .Fnc.camelize .Pkg.Name}}::HttpClient::ResponseHandler.get_body(env)
            message = ""

            # If HTML, whole body is taken
            if body.is_a?(String)
              message = body
            end
{{if .Api.Response.Formats.Json}}
            # If JSON, a particular field is taken and used
            if type && type.include?("json") && body.is_a?(Hash)
              if body.has_key?("{{.Api.Error.Message}}")
                message = body["{{.Api.Error.Message}}"]
              else
                message = "Unable to select error message from json returned by request responsible for error"
              end
            end
{{end}}
            if message == ""
              message = "Unable to understand the content type of response returned by request responsible for error"
            end

            raise {{call .Fnc.camelize .Pkg.Name}}::Error::ClientError.new message, code
          end
        end
      end

    end

  end

end
