package alpaca

import (
	"github.com/robertkrimen/terst"
	"testing"
)

func TestArrayInterfaceToString(t *testing.T) {
	terst.Terst(t)

	var nul interface{}

	terst.Is(ArrayInterfaceToString(nul), []string{})
}

func TestArrayInterfaceInterfaceToString(t *testing.T) {
	terst.Terst(t)

	var nul interface{}

	terst.Is(ArrayInterfaceInterfaceToString(nul, "stuff"), []string{})
}

func TestMapKeysToStringArray(t *testing.T) {
	terst.Terst(t)

	var nul interface{}

	old := make(map[string]interface{})

	old["a"] = "aaa"
	old["b"] = "bbb"
	old["c"] = "ccc"

	terst.Is(MapKeysToStringArray(nul, []string{}), []string{})
	terst.Is(MapKeysToStringArray(old, []string{"b"}), []string{"a", "c"})
}

func TestActiveClassInfo(t *testing.T) {
	terst.Terst(t)

	old := make(map[string]interface{})

	old["args"] = []string{"id", "url"}
	old["create"] = make(map[string]interface{})
	old["update"] = make(map[string]interface{})

	v := ActiveClassInfo("test", old)

	terst.Is(v["name"], "test")
	terst.Is(v["args"].([]string), []string{"id", "url"})
	terst.Is(v["methods"].([]string), []string{"create", "update"})
}

func TestArgsFunctionMaker(t *testing.T) {
	terst.Terst(t)

	var nul interface{}

	f := ArgsFunctionMaker("$", ", ").(func(interface{}, ...bool) string)

	args := make([]interface{}, 2)
	params := make([]interface{}, 2)

	args[0] = "id"
	args[1] = "url"

	params[0] = make(map[string]interface{})
	params[1] = make(map[string]interface{})

	params[0].(map[string]interface{})["name"] = "id"
	params[0].(map[string]interface{})["required"] = true
	params[1].(map[string]interface{})["name"] = "url"

	terst.Is(f(args), "$id, $url, ")
	terst.Is(f(args, false), "$id, $url, ")
	terst.Is(f(args, true), "$id, $url")
	terst.Is(f(args, true, false), "$id, $url")
	terst.Is(f(args, true, true), ", $id, $url")
	terst.Is(f(args, false, true), ", $id, $url, ")

	terst.Is(f(params), "$id, ")
	terst.Is(f(params, true), "$id")
	terst.Is(f(params, false, true), ", $id, ")

	terst.Is(f(nul), "")
	terst.Is(f(nul, true), "")
	terst.Is(f(nul, false, true), "")
}

func TestPathFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := PathFunctionMaker("\"+", "@", "+\"").(func(string, interface{}, interface{}) string)

	cargs := make([]interface{}, 2)
	margs := make([]interface{}, 1)

	cargs[0] = "id"
	cargs[1] = "url"

	margs[0] = make(map[string]interface{})
	margs[0].(map[string]interface{})["name"] = "one"

	terst.Is(f("/user/:id/:not/:url/wow:one", cargs, margs), "/user/\"+@id+\"/:not/\"+@url+\"/wow\"+one+\"")
}

func TestPrntFunctionMaker(t *testing.T) {
	terst.Terst(t)

	var nul interface{}

	f := PrntFunctionMaker(true, "  ", "'", "'", "[", "]", "{", "}", ":", " => ").(func(interface{}, string, bool) string)

	args := make(map[string]interface{})
	vals := make(map[string]interface{})
	orgs := make([]interface{}, 3)
	hash := make(map[string]interface{})

	orgs[0] = false
	orgs[1] = "alpaca-api"
	orgs[2] = 00

	hash["html"] = "haml"
	hash["rest"] = false

	vals["msid"] = 3737
	vals["plan"] = 1.99
	vals["name"] = "pksunkara"
	vals["hire"] = true
	vals["orgs"] = orgs
	vals["hash"] = hash
	vals["dump"] = make(map[string]string)

	args["id"] = make(map[string]interface{})
	args["id"].(map[string]interface{})["value"] = vals

	terst.Is(f(args, ", ", true), "{\n  :msid => 3737,\n  :plan => 1.99,\n  :name => 'pksunkara',\n  :hire => True,\n  :orgs => [\n    False,\n    'alpaca-api',\n    0\n  ],\n  :hash => {\n    :html => 'haml',\n    :rest => False\n  },\n  :dump => \n}, ")
	terst.Is(f(args, ", ", false), "{\n  :msid => 3737,\n  :plan => 1.99,\n  :name => 'pksunkara',\n  :hire => True,\n  :orgs => [\n    False,\n    'alpaca-api',\n    0\n  ],\n  :hash => {\n    :html => 'haml',\n    :rest => False\n  },\n  :dump => \n}")
	terst.Is(f(nul, ", ", true), "")
}
