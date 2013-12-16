package alpaca

import (
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	TemplateDir string
	LibraryRoot string
)

//TODO: Remove 2nd argument
func TemplateInit(lib string, tmp string) {
	var err error

	LibraryRoot, err = filepath.Abs(lib)
	HandleError(err)

	TemplateDir, err = filepath.Abs(tmp)
	HandleError(err)
}

func ReadTemplate(name string) *template.Template {
	temp, err := template.ParseFiles(path.Clean(name))
	HandleError(err)

	return temp
}

func WriteTemplate(temp *template.Template, out string, data interface{}) {
	file, err := os.Create(path.Clean(out))
	defer file.Close()
	HandleError(err)

	HandleError(temp.Execute(file, data))
}

func ChooseTemplate(name string) func(string, string, interface{}) {
	tmp := path.Join(TemplateDir, name)

	return func(name string, out string, data interface{}) {
		temp := ReadTemplate(path.Join(tmp, name))
		WriteTemplate(temp, out, data)
	}
}
