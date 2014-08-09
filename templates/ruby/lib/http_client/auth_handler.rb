require "base64"

module {{call .Fnc.camelize .Pkg.Name}}

  module HttpClient

    # AuthHandler takes care of devising the auth type and using it
    class AuthHandler < Faraday::Middleware
{{if .Api.Authorization.Basic}}
      HTTP_PASSWORD = 0
{{end}}{{if .Api.Authorization.Header}}
      HTTP_HEADER = 1
{{end}}{{if .Api.Authorization.Oauth}}
      URL_SECRET = 2
      URL_TOKEN = 3
{{end}}
      def initialize(app, auth = {}, options = {})
        @auth = auth
        super(app)
      end

      def call(env)
        if !@auth.empty?
          auth = get_auth_type
          flag = false
{{if .Api.Authorization.Basic}}
          if auth == HTTP_PASSWORD
            env = http_password(env)
            flag = true
          end
{{end}}{{if .Api.Authorization.Header}}
          if auth == HTTP_HEADER
            env = http_header(env)
            flag = true
          end
{{end}}{{if .Api.Authorization.Oauth}}
          if auth == URL_SECRET
            env = url_secret(env)
            flag = true
          end

          if auth == URL_TOKEN
            env = url_token(env)
            flag = true
          end
{{end}}
          if !flag
            raise StandardError.new "Unable to calculate authorization method. Please check"
          end{{if .Api.Authorization.NeedAuth}}
        else
          raise StandardError.new "Server requires authentication to proceed further. Please check"{{end}}
        end

        @app.call(env)
      end

      # Calculating the Authentication Type
      def get_auth_type()
{{if .Api.Authorization.Basic}}
        if @auth.has_key?(:username) and @auth.has_key?(:password)
          return HTTP_PASSWORD
        end
{{end}}{{if .Api.Authorization.Header}}
        if @auth.has_key?(:http_header)
          return HTTP_HEADER
        end
{{end}}{{if .Api.Authorization.Oauth}}
        if @auth.has_key?(:client_id) and @auth.has_key?(:client_secret)
          return URL_SECRET
        end

        if @auth.has_key?(:access_token)
          return URL_TOKEN
        end
{{end}}
        return -1
      end
{{if .Api.Authorization.Basic}}
      # Basic Authorization with username and password
      def http_password(env)
        code = Base64.encode64 "#{@auth[:username]}:#{@auth[:password]}"

        env[:request_headers]["Authorization"] = "Basic #{code}"

        return env
      end
{{end}}{{if .Api.Authorization.Header}}
      # Authorization with HTTP header
      def http_header(env)
        env[:request_headers]["Authorization"] = "{{or .Api.Authorization.HeaderPrefix "token"}} #{@auth[:http_header]}"

        return env
      end
{{end}}{{if .Api.Authorization.Oauth}}
      # OAUTH2 Authorization with client secret
      def url_secret(env)
        query = {
          :client_id => @auth[:client_id],
          :client_secret => @auth[:client_secret]
        }

        merge_query(env, query)
      end

      # OAUTH2 Authorization with access token
      def url_token(env)
        query = { :access_token => @auth[:access_token] }

        merge_query(env, query)
      end
{{end}}
      def query_params(url)
        if url.query.nil? or url.query.empty?
          {}
        else
          Faraday::Utils.parse_query(url.query)
        end
      end

      def merge_query(env, query)
        query = query.update query_params(env[:url])

        env[:url].query = Faraday::Utils.build_query(query)

        return env
      end
    end

  end

end
