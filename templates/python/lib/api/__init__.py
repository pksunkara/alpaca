# Import all the classes into api module{{with $data := .}}{{range .Api.Classes}}
from . import {{call $data.Fnc.underscore .Name}}{{end}}{{end}}
