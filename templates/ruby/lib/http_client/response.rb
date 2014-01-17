module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # Response object contains the response returned by the client
    class Response

      attr_accessor :body, :code, :headers

      def initialize(body, code, headers)
        @body = body
        @code = code
        @headers = headers
      end

    end

  end

end
