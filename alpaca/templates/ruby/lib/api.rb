module {{.Pkg.name}}

  module Api

    class {{call .Fnc.camelize .Api.active.name}}

      def initialize(client)
        @client = client
      end

    end

  end

end
