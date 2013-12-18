lib = File.expand_path('../test/ruby/lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require "test-alpaca"

client = Test::Client.new 
client.client_options.base_default

client = Test::Client.new({}, { :base => 'http://localhost:3001/useless' })

client.client_options.base_default
