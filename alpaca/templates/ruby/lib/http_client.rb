require "{{call .Fnc.underscore .Pkg.name}}/http_client/auth_handler"
require "{{call .Fnc.underscore .Pkg.name}}/http_client/error_handler"
require "{{call .Fnc.underscore .Pkg.name}}/http_client/request_handler"
require "{{call .Fnc.underscore .Pkg.name}}/http_client/response_handler"

module {{.Pkg.name}}

  module HttpClient

    # Main HttpClient which is used by Api classes
    class HttpClient

      attr_accessor :options, :headers

      @headers = {}

      def initialize(auth = {}, options = {})
{{if .Api.authorization.oauth}}
        if auth.is_a? String
          auth = { access_token: auth }
        end
{{end}}
        @options = {
          base: "{{.Api.base}}",{{with .Api.version}}
          api_version: "{{.}}",{{end}}
          user_agent: "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)"
        }

        @options.update options

        @headers = {}

        @headers[:user_agent] = @options[:user_agent]

        if @options.has_key? :headers
          @headers.update @options[:headers]
          @options.delete :headers
        end

        @client = Faraday.new url: @options[:base]
      end

      def get(path, params = {}, options = {})
        request path, nil, "get", options.merge({ query: params })
      end

      def post(path, body = {}, options = {})
        request path, body, "post", options
      end

      def patch(path, body = {}, options = {})
        request path, body, "patch", options
      end

      def delete(path, body = {}, options = {})
        request path, body, "delete", options
      end

      def put(path, body = {}, options = {})
        request path, body, "put", options
      end

      # Intermediate function which does three main things
      #
      # - Transforms the body of request into correct format
      # - Creates the requests with give parameters
      # - Returns response body after parsing it into correct format
      def request(path, body, method, options)
        headers = {}

        options = @options.merge options

        if options.has_key? :headers
          headers = options[:headers]
          options.delete :headers
        end

        headers = @headers.merge headers

        options.delete :body

        body, headers = set_body body, headers, options

        response = create_request method, path, body, headers, options

        body = get_body response
      end

      # Creating a request with the given arguments
      #
      # If api_version is set, appends it immediately after host
      def create_request(method, path, body, headers, options)
        version = options.has_key?(:api_version) ? "/#{options[:api_version]}" : ""
{{if .Api.response.suffix}}
        # Adds a suffix (ex: ".html", ".json") to url
        suffix = options.has_key?(:response_type) ? options[:response_type] : "{{or .Api.response.formats.default "html"}}"
        path = "#{path}.#{suffix}"
{{end}}
        path = "#{version}#{path}"

        instance_eval <<-RUBY, __FILE__, __LINE__ + 1
          @client.#{method} path, body, headers
        RUBY
      end

      # Get response body in correct format
      def get_body(response)
        {{.Pkg.name}}::HttpClient::ResponseHandler.get_body response
      end

      # Set request body in correct format
      def set_body(body, headers, options)
        {{.Pkg.name}}::HttpClient::RequestHandler.set_body body, headers, options
      end

    end

  end

end
