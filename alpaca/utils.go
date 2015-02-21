package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"encoding/json"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
)

func ReadJSON(name string, v interface{}) {
	file, err := os.Open(path.Join(LibraryRoot, name))
	HandleError(err)

	HandleError(json.NewDecoder(file).Decode(v))
}

func MakeDir(name string) {
	HandleError(os.Mkdir(name, 0755))
	MoveDir(name)
}

func MoveDir(name string) {
	HandleError(os.Chdir(name))
}

func ArgsFunctionMaker(before, after string) interface{} {
	return func(args interface{}, options ...bool) string {
		str := ""

		if reflect.TypeOf(args).String() == "[]string" {
			for _, v := range args.([]string) {
				str += before + v + after
			}
		} else {
			for _, v := range args.([]ApiParam) {
				if v.Required {
					str += before + v.Name + after
				}
			}
		}

		if str != "" {
			if len(options) > 0 && options[0] {
				str = str[0 : len(str)-len(after)]
			}

			if len(options) > 1 && options[1] {
				str = after + str
			}
		}

		return str
	}
}

func PathFunctionMaker(before, this, after string) interface{} {
	return func(path string, cargs []string, margs []ApiParam) string {
		for _, v := range cargs {
			reg := regexp.MustCompile(":(" + v + ")")
			path = reg.ReplaceAllString(path, before+this+"$1"+after)
		}

		for _, v := range margs {
			reg := regexp.MustCompile(":(" + v.Name + ")")
			path = reg.ReplaceAllString(path, before+"$1"+after)
		}

		return path
	}
}

func PrntFunctionMaker(boolcap bool, tab, strbeg, strend, arrbeg, arrend, objbeg, objend, keybeg, keyend string) interface{} {
	var vals func(interface{}, ...int) string
	var tabs func(int) string

	arrmid, objmid, newline := ",", ",", "\n"

	tabs = func(level int) string {
		str := ""

		for i := 0; i < level; i++ {
			str += tab
		}

		return str
	}

	vals = func(data interface{}, level ...int) string {
		typ, lev := reflect.TypeOf(data).String(), 1

		if len(level) == 1 {
			lev = level[0]
		}

		if typ == "bool" {
			str := strconv.FormatBool(data.(bool))

			if boolcap {
				str = inflect.Capitalize(str)
			}

			return str
		}

		if typ == "string" {
			return strbeg + data.(string) + strend
		}

		if typ == "int" {
			return strconv.Itoa(data.(int))
		}

		if typ == "float64" {
			return strconv.FormatFloat(data.(float64), 'f', -1, 64)
		}

		if typ == "[]interface {}" {
			str := arrbeg

			for _, v := range data.([]interface{}) {
				str += newline + tabs(lev) + vals(v, lev+1) + arrmid
			}

			return str[0:len(str)-len(arrmid)] + newline + tabs(lev-1) + arrend
		}

		if typ == "map[string]interface {}" {
			str := objbeg

			for k, v := range data.(map[string]interface{}) {
				str += newline + tabs(lev) + keybeg + k + keyend + vals(v, lev+1) + objmid
			}

			return str[0:len(str)-len(objmid)] + newline + tabs(lev-1) + objend
		}

		return ""
	}

	return func(api interface{}, doc map[string]DocParam, sep string, notLast bool) string {
		str, typ := "", reflect.TypeOf(api).String()

		if typ == "[]string" {
			for _, v := range api.([]string) {
				str += vals(doc[v].Value) + sep
			}
		} else if typ == "[]alpaca.ApiParam" {
			for _, v := range api.([]ApiParam) {
				if v.Required {
					str += vals(doc[v.Name].Value) + sep
				}
			}
		} else {
			return str
		}

		if str != "" && !notLast {
			return str[0 : len(str)-len(sep)]
		}

		return str
	}
}
