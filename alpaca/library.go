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

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	HandleError(os.RemoveAll(name))
	MakeDir(name)
}
