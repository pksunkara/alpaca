lib = File.expand_path('../test/ruby/lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require "test-alpaca"

# Client Options

client = Test::Client.new
client.client_options.base

Test::Client.new({}, {
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    'custom-header' => 'custom'
  }
}).client_options.base

# Request Options

client.request_options.base({
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    'custom-header' => 'custom',
    'user-agent' => 'testing again'
  }
})
client.request_options.suffix :response_type => 'png'

# GET request

client.get.api('foo', 'bar')
client.get.options :query => { :foo => 'bar' }

# Responses

response = client.response.basic

client.test.equal :query => {
  :expected => response.code,
  :actual => 206,
  :name => 'The status code is correctly propogated'
}

client.test.equal :query => {
  :expected => response.body,
  :actual => '/',
  :name => 'The response body is correctly propogated'
}

client.test.equal :query => {
  :expected => client.response.header.headers['awesome'],
  :actual => 'wow nice',
  :name => 'The response headers are correctly propogated'
}

client.test.equal :query => {
  :expected => client.response.html.body,
  :actual => 'checking html',
  :name => 'The response body in HTML format is correctly parsed'
}

client.test.equal :query => {
  :expected => client.response.json.body['message'],
  :actual => 'checking json',
  :name => 'The response body in JSON format is correctly parsed'
}

# POST Request

client.post.empty_raw
client.post.options_raw :body => 'hello world'

client.post.empty_form :request_type => 'form'
client.post.api_form('foo', 'bar', { :request_type => 'form' })
client.post.options_form({
  :request_type => 'form',
  :body => { :foo => ['bar', 'baz'] }
})

client.post.empty_json :request_type => 'json'
client.post.api_json('foo', 'bar', { :request_type => 'json' })
client.post.options_json({
  :request_type => 'json',
  :body => { :foo => ['bar', 'baz'] }
})

# HTTP Methods
client.methods.patch
client.methods.put
client.methods.delete

# End tests

client.test.end
