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

type Data struct {
	Pkg map[string]interface{}
	Api map[string]interface{}
	Doc map[string]interface{}
	Fnc map[string]interface{}
}

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

func ChooseTemplate(name string) func(string, string, interface{}) {
	tmp := path.Join(TemplateDir, name)

	return func(name string, out string, data interface{}) {
		temp := ReadTemplate(path.Join(tmp, name))
		WriteTemplate(temp, out, data)
	}
}

func MakeLibraryDir(name string) {
	name = path.Join(LibraryRoot, name)

	//TODO: Should we delete previous code?
	HandleError(os.RemoveAll(name))
	MakeDir(name)
}

func MakeDir(name string) {
	HandleError(os.Mkdir(name, 0755))
	HandleError(os.Chdir(name))
}

func MoveDir(name string) {
	HandleError(os.Chdir(name))
}
