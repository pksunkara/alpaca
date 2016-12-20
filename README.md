# alpaca [![Build Status](https://travis-ci.org/pksunkara/alpaca.svg?branch=master)](https://travis-ci.org/pksunkara/alpaca) [![Gitter chat](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/pksunkara/alpaca)

Api Libraries Powered And Created by Alpaca

---

Tired of maintaining API libraries in different languages for your website API? _This is for you_

You have an API for your website but no API libraries for whatever reason? _This is for you_

You are planning to build an API for your website and develop API libraries? _This is for you_

---

You define your API according to the format given below, __alpaca__ builds the API libraries along with their documentation. All you have to do is publishing them to their respective package managers.

Join us at [gitter](https://gitter.im/pksunkara/alpaca) if you need any help.

## Installation

You can download the binaries (v0.2.1)

 * Architecture i386 [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_linux_386.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_windows_386.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_darwin_386.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_freebsd_386.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_openbsd_386.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_netbsd_386.zip?direct) / [plan9](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_plan9_386.zip?direct) ]
 * Architecture amd64 [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_linux_amd64.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_windows_amd64.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_darwin_amd64.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_freebsd_amd64.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_openbsd_amd64.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_netbsd_amd64.zip?direct) ]
 * Architecture arm [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_linux_arm.tar.gz?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_freebsd_arm.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_netbsd_arm.zip?direct) ]

Or by using deb packages (v0.2.1)

 * [ [i386](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_i386.deb?direct) / [amd64](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_amd64.deb?direct) / [armhf](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.1_armhf.deb?direct) ]

Or by using golang (v1.2)

```bash
# Clone the project into your golang workspace
$ git clone git://github.com/pksunkara/alpaca
$ cd alpaca

# Install program
$ make && make install
```

## Examples

You can find some api definitions in the [examples](https://github.com/pksunkara/alpaca/tree/master/examples) directory. The api libraries generated are at [https://github.com/alpaca-api](https://github.com/alpaca-api)

Completed api definitions are [buffer](https://github.com/pksunkara/alpaca/tree/master/examples/buffer).

## Usage

```bash
$ alpaca </path/to/dir>
```

The path here should be a directory with `api.json`, `pkg.json`, `doc.json`

```
Usage:
  alpaca [options] <dir>

Application Options:
  -v, --version    Show version information

Language Options:
      --no-php     Do not write php library
      --no-python  Do not write python library
      --no-ruby    Do not write ruby library
      --no-node    Do not write node library

Help Options:
  -h, --help       Show this help message
```

_Please remove the comments when actually using these json files_

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
  "base_as_arg": true, // Force Base URL to be an argument in generated clients [optional] (default: false)
  "version": "v1", // Default version for the api (https://api.example.com{/version}/users) [optional]
  "no_verify_ssl": true, // Do not verify SSL cert [optional] (default: false)
  "authorization": { // Authorization strategies
    "need_auth": true, // Authentication is compulsory [optional] (default: false)
    "basic" : true, // Basic authentication [optional] (default: false)
    "header": true, // Token in authorization header [optional] (default: false)
    "header_prefix": "oompaloompa", // The first word in header if using token auth header [optional] (default: token)
    "oauth" : true // OAUTH authorization [optional] (default: false)
  },
  "request": { // Settings for requests to the api
    "formats": { // Format of the request body
      "default": "form", // Default format for the request body [optional] (default: raw)
      "form": true, // Support form-url-encoded? [optional] (default: false)
      "json": true // Support json? [optional] (default: false)
    }
  },
  "response": { // Settings for responses from the api
    "suffix": true, // Should the urls be suffixed with response format? [optional] (default: false)
    "formats": { // Format of the response body
      "default": "json", // Default response format. Used when 'suffix' is 'true'
      "html": true, // Support html? [optional] (default: false)
      "json": true // Support json? [optional] (default: false)
    }
  },
  "error": { // Required if response format is 'json'
    "message": "error" // The field to be used from the response body for error message
  },
  "class": { // The classes for the api
    "users": { // Name of a class of the api
      "args": ["login"], // Arguments required for the api class [optional]
      "profile": { // Name of a method of the api
        "path": "/users/:login/:type", // URL of the api method
        "method": "post", // HTTP method of the api method [optional] (default: get)
        "params": [ // Parameters for the api method [optional]
          {
            "name": "type", // Name of the parameter
            "required": true, // The parameter will become an argument of api method [optional] (default: false)
            "url_use": true // This parameter is only used to build url [optional] (default: false)
          },
          {
            "name": "bio",
            "required": true
          }
        ]
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
    "args": { // Arguments of the api class
      "login": { // Name of the argument
        "desc": "Username of the user", // Description of the argument
        "value": "pksunkara" // Value of the argument in docs
      }
    },
    "profile": { // Name of a method of the api
      "title": "Edit profile", // Title of the api method
      "desc": "Edit the user's profile", // Description of the api method
      "params": { // Parameter of the api class
        "bio": { // Name of the parameter
          "desc": "Short bio in profile", // Description of the parameter
          "value": "I am awesome!" // Value of the parameter in docs
        },
        "type": {
          "desc": "Circle of the profile",
          "value": "friends"
        }
      }
    }
  }
}
```

### Request formats

Supported request formats are `raw`, `form`, `json`. The format `raw` is always true.

This means, the `body` set in the options when calling an API method will be able to be encoded according to the respective `request_type`

__If set to `raw`, body is not modified at all__

### Response formats

Supported response formats are `html`, `json`.

### Authorization strategies

Supported are `basic`, `header`, `oauth`

### Language Versions

Supported programming language versions are:

| Language |   V   |   E   |   R   |   S   |   I   |   O   |   N   |
|----------|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
| node     | 0.8   | 0.9   | 0.10  | 0.11  | 0.12  |       |       |
| php      | 5.4   | 5.5   |       |       |       |       |       |
| python   | 2.6   | 2.7   | 3.2   | 3.3   |       |       |       |
| ruby     | 1.8.7 | 1.9.1 | 1.9.2 | 1.9.3 | 2.0.0 | 2.1.0 | 2.1.1 |

### Package Managers

| Language | Package Manager                                    |
|----------|----------------------------------------------------|
| node     | [https://npmjs.org](https://npmjs.org)             |
| php      | [https://packagist.org](https://packagist.org)     |
| python   | [https://pypi.python.org](https://pypi.python.org) |
| ruby     | [https://rubygems.org](https://rubygems.org)       |

## Testing

Check [here](https://github.com/pksunkara/alpaca/tree/testing) to learn about testing.

## Contributors

Here is a list of [Contributors](https://github.com/pksunkara/alpaca/contributors)

__I accept pull requests and guarantee a reply back within a day__

### TODO

You get internet points for pull requesting the following features.

##### Responses

 * [Add support for XML](https://github.com/pksunkara/alpaca/issues/36)
 * [Add support for CSV](https://github.com/pksunkara/alpaca/issues/36)

##### Requests

 * HTTP Method Overloading
 * What about file uploads?

##### API

 * Check returned status code
 * Special case for 204:true and 404:false

##### Libraries

 * Pagination support
 * Classes inside classes (so on..)
 * Validations for params/body in api methods
 * Allow customization of errors
 * Tests for libraries (lots and lots of tests)

##### Readme

 * [Optional params available](https://github.com/pksunkara/alpaca/issues/57)
 * Return types of api calls

##### Comments

 * The descriptions should be wrapped
 * Align @param descriptions

##### Formats

 * [Support YAML](https://github.com/pksunkara/alpaca/issues/63)
 * [Support API blueprint](https://github.com/pksunkara/alpaca/issues/56)
 * [Support Swagger](https://github.com/pksunkara/alpaca/issues/61)
 * [Support WADL](https://github.com/pksunkara/alpaca/issues/13)
 * [Support JSON Schema](https://github.com/pksunkara/alpaca/issues/17)
 * [Support RAML](https://github.com/pksunkara/alpaca/issues/54)

##### Languages

 * [Support Java](https://github.com/pksunkara/alpaca/issues/11)
 * [Support Go](https://github.com/pksunkara/alpaca/issues/9)
 * [Support Clojure](https://github.com/pksunkara/alpaca/issues/49)
 * [Support Rust](https://github.com/pksunkara/alpaca/issues/62)
 * [Support Swift](https://github.com/pksunkara/alpaca/issues/64)
 * Support C, C++, Perl, Scala, C#, Erlang, Lua, Haskell, D, Julia, Groovy
 * Build cli tool for APIs (bash? python? go?)

### Support Projects

Alternatively, you can write your own converter from `alpaca.json` to the following

* Convert into API Blueprint
* Convert into Swagger

## License

MIT/X11

## Bug Reports

Report [here](http://github.com/pksunkara/alpaca/issues). __Guaranteed reply within a day__.

## Contact

Pavan Kumar Sunkara (pavan.sss1991@gmail.com)

Follow me on [github](https://github.com/users/follow?target=pksunkara), [twitter](http://twitter.com/pksunkara)


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/pksunkara/alpaca/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

