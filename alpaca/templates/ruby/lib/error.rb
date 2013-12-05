require 'error/client_error'

module {{.Pkg.name}}

  module Error

    include {{.Pkg.name}}::Error::ClientError

  end

end
