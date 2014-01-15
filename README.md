# alpaca [![Build Status](https://drone.io/github.com/pksunkara/alpaca/status.png)](https://drone.io/github.com/pksunkara/alpaca/latest)

Api Libraries Powered And Created by Alpaca

---

Tired of maintaining API libraries in different languages for your website API? _This is for you_

You have an API for your website but no API libraries for whatever reason? _This is for you_

You are planning to build an API for your website and develop API libraries? _This is for you_

---

You define your API according to the format given below, __alpaca__ builds the API libraries along with their documentation. All you have to do is publishing them to their respective package managers.

Join us at [gitter](https://gitter.im/pksunkara/alpaca) if you need any help. Or at `#alpaca` on freenode IRC.

## Installation

You can download the binaries (v0.2.0)

 * Architecture i386 [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_linux_386.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_windows_386.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_darwin_386.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_freebsd_386.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_openbsd_386.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_netbsd_386.zip?direct) / [plan9](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_plan9_386.zip?direct) ]
 * Architecture amd64 [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_linux_amd64.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_windows_amd64.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_darwin_amd64.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_freebsd_amd64.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_openbsd_amd64.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_netbsd_amd64.zip?direct) ]
 * Architecture arm [ [linux](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_linux_arm.tar.gz?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_freebsd_arm.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_netbsd_arm.zip?direct) ]

Or by using deb packages (v0.2.0)

 * [ [i386](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_i386.deb?direct) / [amd64](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_amd64.deb?direct) / [armhf](https://dl.bintray.com//content/pksunkara/utils/alpaca_0.2.0_armhf.deb?direct) ]

Or by using golang

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
  "version": "v1", // Default version for the api (https://api.example.com{/version}/users) [optional]
  "no_verify_ssl": true, // Do not verify SSL cert [optional] (default: false)
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
        "params": [ // Parameters for the api method [optional]
          {
            "name": "bio", // Name of the parameter
            "required": true // The parameter will become an argument of api method
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
      "id": { // Name of the argument
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
        }
      }
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

### Language Versions

Supported programming language versions are:

 * __node__: [ 0.8 / 0.9 / 0.10 / 0.11 ]
 * __php__: [ 5.4 / 5.5 ]
 * __python__: [ 2.6 / 2.7 / 3.2 / 3.3 ]
 * __ruby__: [ 1.8.7 / 1.9.1 / 1.9.2 / 1.9.3 / 2.0.0 / 2.1.0 ]

### Package Managers

 * __node__: [https://npmjs.org](https://npmjs.org)
 * __php__: [https://packagist.org](https://packagist.org)
 * __python__: [https://pypi.python.org](https://pypi.python.org)
 * __ruby__: [https://rubygems.org](https://rubygems.org)

## Testing

Check [here](https://github.com/pksunkara/alpaca/tree/testing) to learn about testing.

## Contributors

Here is a list of [Contributors](https://github.com/pksunkara/alpaca/contributors)

__I accept pull requests and guarantee a reply back within a day__

### TODO

##### General

 * Support YAML configuration

##### Responses

 * Add support for XML
 * Add support for CSV

##### Requests

 * HTTP Method Overloading
 * What about file uploads?

##### Api

 * Check returned status code
 * Special case for 204:true and 404:false

##### Libraries

 * Pagination support
 * Classes inside classes (so on..)
 * Validations for params/body in api methods
 * Allow customization of errors
 * Tests for libraries (lots and lots of tests)

##### Readme

 * Optional params available
 * Return types of api calls

##### Comments

 * The descriptions should be wrapped
 * Align @param descriptions

##### Formats

 * Support WADL
 * Support RAML
 * Support API blueprint

##### Languages

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
