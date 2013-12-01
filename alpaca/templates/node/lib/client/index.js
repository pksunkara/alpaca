//var request = require("request");

var client = module.exports;

client.AuthHandler = require("./auth.js");
client.ErrorHandler = require("./error.js")

client.RequestHandler = require("./request.js")
client.ResponseHandler = require("./response.js")

client.HttpClient = function (auth, options) {
{{if .Api.authorization.oauth}}
  if (typeof auth == "string") {
    auth = { "access_token": auth };
  }
{{end}}
  this.options = {
    "base": "{{.Api.base}}",{{with .Api.version}}
    "api_version": "{{.}}",{{end}}
    "user_agent": "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)"
  };

  for (key in options) {
    this.options[key] = options[key];
  }

  this.headers = {
    "User-Agent": this.options["user_agent"]
  };

  if (this.options["headers"]) {
    for (key in this.options["headers"]) {
      this.headers[key] = this.options["headers"][key];
    }

    delete this.options["headers"];
  }

  return this;
}

client.HttpClient.prototype.get = function (path, params, options, callback) {
  if (typeof params == "function") {
    callback = params;
    params = {};
    options = {};
  } else if (typeof options == "function") {
    callback = options;
    options = {};
  }

  options["query"] = params;

  this.request(path, {}, 'GET', options, callback);
};

client.HttpClient.prototype.post = function (path, body, options, callback) {
  if (typeof options == "function") {
    callback = options;
    options = {};
  }

  this.request(path, body, 'POST', options, callback);
};

client.HttpClient.prototype.patch = function (path, body, options, callback) {
  if (typeof options == "function") {
    callback = options;
    options = {};
  }

  this.request(path, body, 'PATCH', options, callback);
};

client.HttpClient.prototype.delete = function (path, body, options, callback) {
  if (typeof options == "function") {
    callback = options;
    options = {};
  }

  this.request(path, body, 'DELETE', options, callback);
};

client.HttpClient.prototype.put = function (path, body, options, callback) {
  if (typeof options == "function") {
    callback = options;
    options = {};
  }

  this.request(path, body, 'PUT', options, callback);
};

client.HttpClient.prototype.request = function (path, body, method, options, callback) {
  headers = {};

  if (options["headers"]) {
    headers = options["headers"];
    delete options["headers"];
  }


};

client.HttpClient.prototype.getBody = function () {

};

client.HttpClient.prototype.createBody = function () {

};
