require "rubygems"

require "{{call .Fnc.underscore .Pkg.Name}}/client"
require "{{call .Fnc.underscore .Pkg.Name}}/error"
require "{{call .Fnc.underscore .Pkg.Name}}/http_client"

module {{call .Fnc.camelize .Pkg.Name}}
end
