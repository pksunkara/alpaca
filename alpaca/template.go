package alpaca

import (
	"github.com/pksunkara/alpaca/templates"
	"os"
	"path"
	"text/template"
)

func ReadTemplate(name string) *template.Template {
	temp := string(templates.Data[name]())

	return template.Must(template.New("temp").Parse(temp))
}

func WriteTemplate(temp *template.Template, out string, data interface{}) {
	file, err := os.Create(path.Clean(out))
	defer file.Close()
	HandleError(err)

	HandleError(temp.Execute(file, data))
}

func ChooseTemplate(template string) func(string, string, interface{}) {
	return func(name string, out string, data interface{}) {
		temp := ReadTemplate(path.Join(template, name))
		WriteTemplate(temp, out, data)
	}
}
