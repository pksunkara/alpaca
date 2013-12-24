lib = File.expand_path('../test/ruby/lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require "test-alpaca"

client = Test::Client.new({}, { :suffix => false })
client.client_options.base

Test::Client.new({}, {
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    'custom-header' => 'custom'
  },
  :suffix => false
}).client_options.base

client.request_options.base({
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    'custom-header' => 'custom'
  }
})

client.methods.get
client.methods.post
client.methods.put
client.methods.patch
client.methods.delete
