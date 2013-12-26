# Import all the classes into api module{{with $data := .}}{{range .Api.classes}}
from . import {{call $data.Fnc.underscore .}}{{end}}{{end}}
