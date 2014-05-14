## 0.2.1

Using Mozilla 2.0 license from now on.

Features:

 - Support `authorization.header_prefix` for authorization headers (#6)
 - Can now make authentication compulsory using `authorizaiton.need_auth`
 - Default value is `false` for `response.formats.html`
 - Default value is `false` for `request.formats.form`
 - Allow API `base_url` to be an argument (#30)
 - Parameters needed for method URL can be defined in the method

Bugfixes:

 - Helpful error when missing language specific fields in `pkg.json`
 - Better building of binary (#22)
 - Python style fixes (#26)
 - Comments in generated code are now params/args aware
 - Fix bug with JSON parsing response in node.js
 - Ruby style fixes

## 0.2.0

Features:

 - Added Makefile (#7)
 - Added support for `no_verify_ssl` in **api.json** (#10)
 - Made `params` in **api.json** into an array of hashes (#12)

Bugfixes:

 - MakeStringArrayInterface and ArrayStringInterface supports nil (#14)
