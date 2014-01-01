# alpaca [![Build Status](https://drone.io/github.com/pksunkara/alpaca/status.png)](https://drone.io/github.com/pksunkara/alpaca/latest)

Api Libraries Powered And Created by Alpaca

---

Tired of maintaining API libraries in different languages for your website API? _This is for you_

You have an API for your website but no API libraries for whatever reason? _This is for you_

You are planning to build an API for your website and develop API libraries? _This is for you_

---

You define your API, **alpaca** builds the API libraries along with their documentation. All you have to do is publishing them to their respective package managers.

Join us at [gitter](https://gitter.im/pksunkara/alpaca) if you need any help. Or at `#alpaca` on freenode IRC.

## Installation

You can download the binary files.

Or by using golang

```bash
# Clone the project into your golang workspace
$ git clone git://github.com/pksunkara/alpaca

# Compile templates
$ go get github.com/jteeuwen/go-bindata
$ cd alpaca && ./make

# Install the program
$ go get
$ go install github.com/pksunkara/alpaca
```

## Usage

```bash
$ alpaca </path/to/dir>
```

The path here should be a directory with `api.json`, `pkg.json`, `doc.json`

#### pkg.json

All the following fields are required unless mentioned.

```js
{
  "name": "Example", // Name of the api (also used as class name for the library)
  "package": "example-alpaca", // Name of the package
  "version": "0.1.0", // Version of the package
  "url": "https://exampleapp.com", // URL of the api
  "keywords": ["alpaca", "exampleapp", "api"], // Keywords for the package
  "official": false, // Are the api libraries official?
  "author": {
    "name": "Pavan Kumar Sunkara", // Name of the package author
    "email": "pavan.sss1991@gmail.com", // Email of the package author
    "url": "http://github.com/pksunkara" // URL of the package author
  },
  "git": { // Used in the package definition
    "site": "github.com", // Name of the git website
    "user": "alpaca-api", // Username of the git website
    "name": "buffer" // Namespace of the git repositories
  },
  "license": "MIT", // License of the package
  "php": { // Required only if creating php api lib
    "vendor": "pksunkara" // Packagist vendor name for the package
  },
  "python": { // Required only if creating python api lib
    "license": "MIT License" // Classifier of the license used for the module
  }
}
```

#### api.json

All the following fields are required unless mentioned.

```js
{
  "base": "https://exampleapp.com", // Base URL of the api
  "version": "v1", // Default version for the api (https://api.example.com{/version}/users) [optional]
  "authorization": { // Authorization strategies
    "basic" : true, // Basic authentication [optional] (default: false)
    "header": true, // Token in authorization header [optional] (default: false)
    "oauth" : true // OAUTH authorization [optional] (default: false)
  },
  "request": { // Settings for requests to the api
    "formats": { // Format of the request body
      "default": "form", // Default format for the request body [optional] (default: raw)
      "json": true // Support json? [optional] (default: false)
    }
  },
  "response": { // Settings for responses from the api
    "formats": { // Format of the response body
      "default": "json", // Default response format. Used when 'suffix' is true [optional] (default: html) 
      "json": true // Support json? [optional] (default: false)
    },
    "suffix": true // Should the urls be suffixed with response format? [optional] (default: false)
  },
  "error": { // Required if response format is 'json'
    "message": "error" // The field to be used from the response body for error message
  },
  "class": { // The classes for the api
    "users": { // Name of a class of the api
      "args": ["login"], // Arguments required for the api class [optional]
      "profile": { // Name of a method of the api
        "path": "/users/:login/profile", // Url of the api method
        "method": "post", // HTTP method of the api method [optional] (default: get)
        "params": ["bio"] // Arguments required for the api method [optional]
      }
    }
  }
}
```

#### doc.json

The following is filled according to the entries in `api.json`

```js
{
  "users": { // Name of a class of the api
    "title": "Users", // Title of the api class
    "desc": "Returns user api instance", // Description of the api class
    "args": [{
      "desc": "Username of the user", // Description of the argument
      "value": "pksunkara" // Value of the argument in docs
    }],
    "profile": { // Name of a method of the api
      "title": "Edit profile", // Title of the api method
      "desc": "Edit the user's profile", // Description of the api method
      "params": [{
        "desc": "Short bio in profile", // Description of the argument
        "value": "I am awesome!" // Value of the argument in docs
      }]
    }
  }
}
```

### Request formats

Supported request formats are `raw`, `form`, `json`.

The formats `raw` and `form` are always true.

### Response formats

Supported response formats are `html`, `json`.

The format `html` is always true.

### Authorization strategies

Supported are `basic`, `header`, `oauth`

## Testing

Check [here](https://github.com/pksunkara/alpaca/tree/testing) to learn about testing.

## Contributors

Here is a list of [Contributors](https://github.com/pksunkara/alpaca/contributors)

__I accept pull requests and guarantee a reply back within a day__

### TODO

#### General

 * Convert `make` into `Makefile`

#### Responses

 * Add support for XML
 * Add support for CSV

#### Requests

 * HTTP Method Overloading
 * What about file uploads?

#### Api

 * Check returned status code
 * Special case for 204:true and 404:false

#### Libraries

 * Pagination support
 * Classes inside classes (so on..)
 * Validations for params/body in api methods
 * Allow customization of errors
 * Tests for libraries (lots and lots of tests)

#### Readme

 * Examples in place of args/params
 * Return types of api calls
 * Options available

#### Comments

 * The descriptions should be wrapped
 * Align @param descriptions

#### Languages

 * Support Java, Go, Perl, Clojure, Scala, Obj-C
 * Build API docs (Resulting in bloated definitions?)
 * Build cli tool for APIs (bash? python? go?)

## License

MIT/X11

## Bug Reports

Report [here](http://github.com/pksunkara/alpaca/issues). __Guaranteed reply within a day__.

## Contact

Pavan Kumar Sunkara (pavan.sss1991@gmail.com)

Follow me on [github](https://github.com/users/follow?target=pksunkara), [twitter](http://twitter.com/pksunkara)
