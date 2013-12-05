# alpaca

Api Libraries Powered And Created by Alpaca

COMING SOON!

---

Tired of maintaining API libraries in different languages for your website API? _This is for you_

You have an API for your website but no API libraries for whatever reason? _This is for you_

You are planning to build an API for your website and develop API libraries? _This is for you_

You define your API, **alpaca** builds the API libraries along with the documentation. All you have to do is publishing them to their respective package managers.

## Installation

You can download the binary files.

## Usage

```
alpaca </path/to/dir>
```

The path here should be a directory with `api.json`, `pkg.json`, `doc.json`

#### pkg.json

All the following fields are required unless mentioned.

```js
{
  "name": "Buffer", // Name of the api (also used as class name for the library)
  "package": "buffer-alpaca", // Name of the package
  "version": "0.1.0", // Version of the package
  "url": "https://bufferapp.com", // URL of the api
  "keywords": ["alpaca", "buffer", "api"], // Keywords for the package
  "official": false, // Is the api official?
  "author": {
    "name": "Pavan Kumar Sunkara", // Name of the package author
    "email": "pavan.sss1991@gmail.com", // Email of the package author
    "url": "http://github.com/pksunkara" // URL of the package author
  },
  "git": {
    "site": "github.com", // Name of the git website
    "user": "alpaca-api", // Username of the git website
    "name": "buffer" // Namespace of the git repositories
  },
  "license": "MIT", // License of the package
  "php": { // Required only if creating php api lib
    "vendor": "pksunkara" // Packagist vendor name for the package
  },
  "python": { // Optional when creating python api lib
    "license": "MIT License" // Classifier of the license used for the module
  }
}
```

## Contributors

Here is a list of [Contributors](http://github.com/pksunkara/alpaca/contributors)

__I accept pull requests and guarantee a reply back within a day__

## TODO

#### Responses
- Add support for XML
- Add support for CSV

#### Requests
- HTTP Method Overloading
- What about file uploads?

#### Api
- Check returned status code
- Output the headers too (somehow)
- Special case for 204:true and 404:false

#### General
- Pagination support
- Classes inside classes (so on..)
- Validations for params/body in api methods
- Allow customization of errors

#### Comments
- The descriptions should be wrapped
- Align @param descriptions

#### Languages
- Support Java, Perl, Clojure & Go
- Build API docs (Resulting in bloated definitions?)
- Build cli tool for APIs (bash? python? go?)

## License

MIT/X11

## Bug Reports

Report [here](http://github.com/pksunkara/alpaca/issues). __Guaranteed reply within a day__.

## Contact

Pavan Kumar Sunkara (pavan.sss1991@gmail.com)

Follow me on [github](https://github.com/users/follow?target=pksunkara), [twitter](http://twitter.com/pksunkara)
