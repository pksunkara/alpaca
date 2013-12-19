lib = File.expand_path('../test/ruby/lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require "test-alpaca"

client = Test::Client.new
client.client_options.base

Test::Client.new({}, {
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    "custom-header" => 'custom'
  }
}).client_options.base

client.methods.get
client.methods.post
