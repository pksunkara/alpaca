package langs

import (
	"regexp"
)

func ArrayInterfaceToString(inter interface{}) []string {
	old := inter.([]interface{})
	new := make([]string, len(old))

	for i, v := range old {
		new[i] = v.(string)
	}

	return new
}

func MapKeysToStringArray(inter interface{}, exclude []string) []string {
	old := inter.(map[string]interface{})
	new := make([]string, 0, len(old))

	for v, _ := range old {
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

func ActiveClassInfo(name string, class interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	data["name"] = name
	data["methods"] = MapKeysToStringArray(class, []string{"args"})
	data["args"] = class.(map[string]interface{})["args"]

	return data
}

func ArgsFunctionMaker(before, after string) interface{} {
	return func(class interface{}, key string, last bool) string {
		str, args := "", class.(map[string]interface{})[key]

		if args != nil {
			for _, v := range ArrayInterfaceToString(args) {
				str += before + v + after
			}
			if last {
				str = str[0 : len(str)-len(after)]
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

func CounterTracker() map[string]interface{} {
	val, fnc := 0, make(map[string]interface{})

	fnc["value"] = func() int {
		val = val + 1
		return val - 1
	}

	fnc["start"] = func() string {
		val = 0
		return ""
	}

	return fnc
}
