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

func ArrayInterfaceToString(inter interface{}) []string {
	if inter == nil {
		return []string{}
	}

	old := inter.([]interface{})
	new := make([]string, len(old))

	for i, v := range old {
		new[i] = v.(string)
	}

	return new
}

func MapKeysToStringArray(inter interface{}, exclude []string) []string {
	if inter == nil {
		return []string{}
	}

	old := inter.(map[string]interface{})
	new := make([]string, 0, len(old))

	for v := range old {
		flag := true

		for _, e := range exclude {
			if e == v {
				flag = false
			}
		}

		if flag {
			new = append(new, v)
		}
	}

	return new
}

func MethodList(class interface{}) []string {
	return MapKeysToStringArray(class, []string{"args"})
}

func ActiveClassInfo(name string, class interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	data["name"] = name
	data["methods"] = MethodList(class)
	data["args"] = class.(map[string]interface{})["args"]

	return data
}

func ArgsFunctionMaker(before, after string) interface{} {
	return func(args interface{}, options ...bool) string {
		str := ""

		if args != nil {
			for _, v := range args.([]interface{}) {
				if reflect.TypeOf(v).String() == "string" {
					str += before + v.(string) + after
				} else {
					val := v.(map[string]interface{})

					if val["required"] != nil && val["required"].(bool) {
						str += before + val["name"].(string) + after
					}
				}
			}

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

func PathFunctionMaker(before, after string) interface{} {
	return func(path string, args interface{}) string {
		if args != nil {
			for _, v := range ArrayInterfaceToString(args) {
				reg := regexp.MustCompile(":(" + v + ")")
				path = reg.ReplaceAllString(path, before+"$1"+after)
			}
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

	return func(args interface{}, sep string, notLast bool) string {
		str := ""

		if args == nil {
			return str
		}

		for _, v := range args.(map[string]interface{}) {
			str += vals(v.(map[string]interface{})["value"]) + sep
		}

		if !notLast {
			return str[0 : len(str)-len(sep)]
		}

		return str
	}
}
