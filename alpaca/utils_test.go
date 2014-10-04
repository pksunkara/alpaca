package alpaca

import (
	"github.com/robertkrimen/terst"
	"testing"
)

func TestArgsFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := ArgsFunctionMaker("$", ", ").(func(interface{}, ...bool) string)

	args := []string{"id", "url"}
	params := make([]ApiParam, 2)

	nulArgs := []string{}
	nulParams := make([]ApiParam, 0)

	params[0] = ApiParam{"id", true, false}
	params[1] = ApiParam{"url", false, false}

	terst.Is(f(args), "$id, $url, ")
	terst.Is(f(args, false), "$id, $url, ")
	terst.Is(f(args, true), "$id, $url")
	terst.Is(f(args, true, false), "$id, $url")
	terst.Is(f(args, true, true), ", $id, $url")
	terst.Is(f(args, false, true), ", $id, $url, ")

	terst.Is(f(params), "$id, ")
	terst.Is(f(params, true), "$id")
	terst.Is(f(params, false, true), ", $id, ")

	terst.Is(f(nulArgs), "")
	terst.Is(f(nulArgs, true), "")
	terst.Is(f(nulArgs, false, true), "")

	terst.Is(f(nulParams), "")
	terst.Is(f(nulParams, true), "")
	terst.Is(f(nulParams, false, true), "")
}

func TestPathFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := PathFunctionMaker("\"+", "@", "+\"").(func(string, []string, []ApiParam) string)

	cargs := []string{"id", "url"}
	margs := make([]ApiParam, 1)

	margs[0] = ApiParam{"one", false, false}

	terst.Is(f("/user/:id/:not/:url/wow:one", cargs, margs), "/user/\"+@id+\"/:not/\"+@url+\"/wow\"+one+\"")
}

func TestPrntFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := PrntFunctionMaker(true, "  ", "'", "'", "[", "]", "{", "}", ":", " => ").(func(map[string]DocParam, string, bool) string)

	args := make(map[string]DocParam)
	vals := make(map[string]interface{})
	orgs := make([]interface{}, 3)
	null := make(map[string]string)

	orgs[0] = false
	orgs[1] = "alpaca-api"
	orgs[2] = 00

	terst.Is(f(make(map[string]DocParam), ", ", true), "")

	vals["key"] = 3737
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", true), "{\n  :key => 3737\n}, ")

	vals["key"] = 1.99
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", false), "{\n  :key => 1.99\n}")

	vals["key"] = "pksunkara"
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", false), "{\n  :key => 'pksunkara'\n}")

	vals["key"] = true
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", false), "{\n  :key => True\n}")

	vals["key"] = orgs
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", false), "{\n  :key => [\n    False,\n    'alpaca-api',\n    0\n  ]\n}")

	vals["key"] = null
	args["id"] = DocParam{"", vals}
	terst.Is(f(args, ", ", false), "{\n  :key => \n}")
}
