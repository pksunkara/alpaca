package langs

import (
	"github.com/robertkrimen/terst"
	"testing"
)

func TestArrayInterfaceToString(t *testing.T) {
	terst.Terst(t)

	old := make([]interface{}, 2)

	old[0] = "hello"
	old[1] = "world"

	terst.Is(ArrayInterfaceToString(old), []string{"hello", "world"})
}

func TestMapKeysToStringArray(t *testing.T) {
	terst.Terst(t)

	old := make(map[string]interface{})

	old["a"] = "aaa"
	old["b"] = "bbb"
	old["c"] = "ccc"

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

	f := ArgsFunctionMaker("$", ", ").(func(interface{}, bool) string)
	args := make([]interface{}, 2)

	args[0] = "id"
	args[1] = "url"

	terst.Is(f(args, false), "$id, $url, ")
	terst.Is(f(args, true), "$id, $url")
}

func TestPathFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := PathFunctionMaker("\"+@", "+\"").(func(string, interface{}) string)
	args := make([]interface{}, 2)

	args[0] = "id"
	args[1] = "url"

	terst.Is(f("/user/:id/:not/:url/wow", args), "/user/\"+@id+\"/:not/\"+@url+\"/wow")
}

func TestCounterTracker(t *testing.T) {
	terst.Terst(t)

	f := CounterTracker()

	terst.Is(f["start"].(func() string)(), "")

	for i := 0; i < 5; i++ {
		terst.Is(f["value"].(func() int)(), i)
	}
}
