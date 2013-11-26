package alpaca

import (
	"os"
	"path"
	"text/template"
)

func MakeTemplate(name string) *template.Template {
	temp, err := template.ParseFiles(path.Clean(name))
	HandleError(err)

	return temp
}

func RunTemplate(name string, data interface{}) {
	temp := MakeTemplate(name)
	HandleError(temp.Execute(os.Stdout, data))
}
