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

	f := PrntFunctionMaker(true, "  ", "'", "'", "[", "]", "{", "}", ":", " => ").(func(interface{}, map[string]DocParam, string, bool) string)

	apis := []string{"id"}
	apip := make([]ApiParam, 2)
	docs := make(map[string]DocParam)
	vals := make(map[string]interface{})
	orgs := make([]interface{}, 3)
	null := make(map[string]string)

	orgs[0] = false
	orgs[1] = "alpaca-api"
	orgs[2] = 00

	apip[0] = ApiParam{"id", true, false}
	apip[1] = ApiParam{"flag", false, false}

	docs["flag"] = DocParam{"", false}

	terst.Is(f([]string{}, make(map[string]DocParam), ", ", true), "")
	terst.Is(f([]ApiParam{}, make(map[string]DocParam), ", ", true), "")
	terst.Is(f([]int{}, make(map[string]DocParam), ", ", true), "")

	vals["key"] = 3737
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", true), "{\n  :key => 3737\n}, ")
	terst.Is(f(apip, docs, ", ", true), "{\n  :key => 3737\n}, ")

	vals["key"] = 1.99
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", false), "{\n  :key => 1.99\n}")
	terst.Is(f(apip, docs, ", ", false), "{\n  :key => 1.99\n}")

	vals["key"] = "pksunkara"
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", false), "{\n  :key => 'pksunkara'\n}")
	terst.Is(f(apip, docs, ", ", false), "{\n  :key => 'pksunkara'\n}")

	vals["key"] = true
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", false), "{\n  :key => True\n}")
	terst.Is(f(apip, docs, ", ", false), "{\n  :key => True\n}")

	vals["key"] = orgs
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", false), "{\n  :key => [\n    False,\n    'alpaca-api',\n    0\n  ]\n}")
	terst.Is(f(apip, docs, ", ", false), "{\n  :key => [\n    False,\n    'alpaca-api',\n    0\n  ]\n}")

	vals["key"] = null
	docs["id"] = DocParam{"", vals}
	terst.Is(f(apis, docs, ", ", false), "{\n  :key => \n}")
	terst.Is(f(apip, docs, ", ", false), "{\n  :key => \n}")
}
