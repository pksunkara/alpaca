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
    'custom-header' => 'custom'
  }
}).client_options.base

client.request_options.base({
  :base => 'http://localhost:3001/useless',
  :api_version => 'v2',
  :user_agent => 'testing (user agent)',
  :headers => {
    'custom-header' => 'custom'
  }
})
client.request_options.suffix :response_type => 'png'

client.get.api_params('foo', 'bar')
client.get.query_options :query => { :foo => 'bar' }

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

client.test.end
