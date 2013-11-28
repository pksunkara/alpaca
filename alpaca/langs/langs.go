package langs

import (
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	HandleError func(error)
	TemplateDir string
	LibraryRoot string
)

//TODO: Remove 3rd argument
func Init(fun func(error), lib string, tmp string) {
	var err error

	HandleError = fun
	LibraryRoot, err = filepath.Abs(lib)
	HandleError(err)
	TemplateDir, _ = filepath.Abs(tmp)
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

func RunTemplate(name string, out string, data interface{}) {
	temp := ReadTemplate(path.Join(TemplateDir, name))
	WriteTemplate(temp, out, data)
}

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	//TODO: Should we delete previous code?
	HandleError(os.RemoveAll(name))
	HandleError(os.Mkdir(name, 0755))
	HandleError(os.Chdir(name))
}
