package alpaca

import (
	"os"
	"path"
)

type PkgStruct struct {
	Name    string
	Package string
	Version string
	Url     string

	Keywords []string
	Official bool
	License  string

	Author struct {
		Name  string
		Email string
		Url   string
	}

	Git struct {
		Site string
		User string
		Name string
	}

	Php struct {
		Vendor string
	}

	Python struct {
		License string
	}
}

type ApiParam struct {
	Name     string
	Required bool
	UrlUse   bool `json:"url_use"`
}

type ApiFunction struct {
	Name   string
	Path   string
	Method string

	Params []ApiParam
}

type ApiClass struct {
	Name string
	Args []string

	Functions []ApiFunction
}

type ApiStruct struct {
	Version string
	Base    string

	BaseAsArg   bool `json:"base_as_arg"`
	NoVerifySSL bool `json:"no_verify_ssl"`

	Authorization struct {
		Basic  bool
		Oauth  bool
		Header bool

		HeaderPrefix string `json:"header_prefix"`
		NeedAuth     bool   `json:"need_auth"`
	}

	Request struct {
		Formats struct {
			Default string

			Form bool
			Json bool
		}
	}

	Response struct {
		Suffix bool

		Formats struct {
			Default string

			Html bool
			Json bool
		}
	}

	Error struct {
		Message string
	}

	Classes []ApiClass
}

type DocParam struct {
	Desc  string
	Value interface{}
}

type DocFunction struct {
	Title string
	Desc  string

	Params map[string]DocParam
}

type DocClass struct {
	Title string
	Desc  string

	Args      map[string]DocParam
	Functions map[string]DocFunction
}

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	HandleError(os.RemoveAll(name))
	MakeDir(name)
}
