require "base64"

module {{.Pkg.name}}

  module HttpClient

    # AuthHandler takes care of devising the auth type and using it
    class AuthHandler < Faraday::Middleware

      HTTP_PASSWORD = 0
      HTTP_TOKEN = 1
{{if .Api.authorization.oauth}}
      URL_SECRET = 2
      URL_TOKEN = 3
{{end}}
      def initialize(app, auth = {}, options = {})
        @auth = auth
        super(app)
      end

      def call(env)
        if !@auth.empty?
          case get_auth_type

          when HTTP_PASSWORD
            env = http_password env
          when HTTP_TOKEN
            env = http_token env{{if .Api.authorization.oauth}}
          when URL_SECRET
            env = url_secret env
          when URL_TOKEN
            env = url_token env{{end}}
          else
            raise StandardError.new "Unable to calculate authorization method. Please check"
          end
        end

        @app.call(env)
      end

      # Calculating the Authentication Type
      def get_auth_type()
        if @auth.has_key?(:username) and @auth.has_key?(:password)
          return HTTP_PASSWORD
        elsif @auth.has_key?(:http_token)
          return HTTP_TOKEN{{if .Api.authorization.oauth}}
        elsif @auth.has_key?(:client_id) and @auth.has_key?(:client_secret)
          return URL_SECRET
        elsif @auth.has_key?(:access_token)
          return URL_TOKEN{{end}}
        else
          return -1
        end
      end

      # Basic Authorization with username and password
      def http_password(env)
        code = Base64.encode64 "#{@auth[:username]}:#{@auth[:password]}"

        env[:headers]["Authorization"] = "Basic #{code}"

        return env
      end

      # Authorization with HTTP token
      def http_token(env)
        env[:headers]["Authorization"] = "token #{@auth[:http_token]}"

        return env
      end
{{if .Api.authorization.oauth}}
      # OAUTH2 Authorization with client secret
      def url_secret(env)
        query = {
          client_id: @auth[:client_id],
          client_secret: @auth[:client_secret]
        }

        merge_query env, query
      end

      # OAUTH2 Authorization with access token
      def url_token(env)
        query = { access_token: @auth[:access_token] }

        merge_query env, query
      end
{{end}}
      def query_params(url)
        if url.query.nil? or url.query.empty?
          {}
        else
          Faraday::Utils.parse_query url.query
        end
      end

      def merge_query(env, query)
        query = query.update query_params(env[:url])

        env[:url].query = Faraday::Utils.build_query query

        return env
      end
    end

  end

end
