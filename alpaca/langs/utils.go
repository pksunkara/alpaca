package langs

func ArrayInterfaceToString(inter interface{}) []string {
	old := inter.([]interface{})
	new := make([]string, len(old))

	for i, v := range old {
		new[i] = v.(string)
	}

	return new
}

func MapKeysToStringArray(inter interface{}) []string {
	old := inter.(map[string]interface{})
	new := make([]string, 0, len(old))

	for v, _ := range old {
		new = append(new, v)
	}

	return new
}

func ArgsTemplate(before string, after string) interface{} {
	return func(class interface{}, last bool) string {
		str, args := "", class.(map[string]interface{})["args"]

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
