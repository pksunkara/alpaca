package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/wsxiaoys/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	ClassCount int
	SuiteCount int
	TestsCount int
)

func main() {
	app := createServer()
	tmp := createServer()

	ClassCount = 0
	SuiteCount = 0
	TestsCount = 0

	app.Get("/v1/index.json", func(rw http.ResponseWriter, req *http.Request) {
		Suite("Client Options", 8)
		Test("Base value given by api works correctly")
		Test("Api version value given by api works correctly")

		if req.UserAgent() == "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)" {
			Test("User agent default value works correctly")
		}

		fmt.Fprintf(rw, "/")
	})

	tmp.Get("/v2/index.json", func(rw http.ResponseWriter, req *http.Request) {
		Test("Base value given by options works correctly")
		Test("Api version value given by options works correctly")

		if req.UserAgent() == "testing (user agent)" {
			Test("User agent value given by options works correctly")
		}

		if req.Header["Custom-Header"][0] == "custom" {
			Test("Headers value given by options works correctly")
		}

		Test("Base value's path should not be used")

		fmt.Fprintf(rw, "/")
	})

	app.Get("/v2/index.json", func(rw http.ResponseWriter, req *http.Request) {
		Suite("Request Options", 5)
		Test("Base value should not be changed")
		Test("Api version value given by options works correctly")

		if req.Header["Custom-Header"][0] == "custom" {
			Test("Headers value given by options works correctly")
		}

		if req.UserAgent() == "testing again" {
			Test("Headers value is merged with client options headers correctly")
		}

		fmt.Fprintf(rw, "/")
	})

	app.Get("/v1/suffix.png", func() string {
		Test("Response type used for suffix given by options")
		return "/"
	})

	app.Get("/v1/test/equal.json", func(rw http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()

		if values.Get("expected") == values.Get("actual") {
			Test(values.Get("name"))
		}
	})

	app.Get("/v1/test/end.json", func(rw http.ResponseWriter, req *http.Request) {
		num := 8

		if ClassCount != num {
			terminal.Stdout.Nl().Color("r!").Print("Missing " + strconv.Itoa(num-ClassCount) + " sectons of tests").Nl()
		}

		go os.Exit(0)
	})

	app.Get("/v1/get/api.json", func(rw http.ResponseWriter, req *http.Request) {
		Suite("GET Request", 3)
		Test("Basic request is successful")

		if req.URL.RawQuery == "first=foo&second=bar" {
			Test("Query params using api works correctly")
		}
	})

	app.Get("/v1/get/options.json", func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.RawQuery == "foo=bar" {
			Test("Query params using options works correctly")
		}
	})

	app.Get("/v1/response/basic.json", func() (int, string) {
		Suite("Responses", 5)
		return 206, "/"
	})

	app.Get("/v1/response/header.json", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("awesome", "wow nice")
	})

	app.Get("/v1/response/html.json", func() string {
		return "checking html"
	})

	app.Get("/v1/response/json.json", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("content-type", "application/json; charset=utf-8")
		fmt.Fprintf(rw, "{\"message\": \"checking json\"}")
	})

	app.Post("/v1/post/empty_raw.json", func(rw http.ResponseWriter, req *http.Request) {
		Suite("POST Request", 9)
		Test("Basic request is successful")

		if data, _ := ioutil.ReadAll(req.Body); string(data) == "" {
			Test("Empty raw body works correctly")
		}
	})

	app.Post("/v1/post/options_raw.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "hello world" {
			Test("Setting raw body using options works correctly")
		}
	})

	app.Post("/v1/post/empty_form.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "" {
			Test("Empty form body works correctly")
		}
	})

	app.Post("/v1/post/api_form.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "first=foo&second=bar" {
			Test("Setting form body using api works correctly")
		}
	})

	app.Post("/v1/post/options_form.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "foo%5B%5D=bar&foo%5B%5D=baz" {
			Test("Setting form body using options works correctly")
		}
	})

	app.Post("/v1/post/empty_json.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "{}" {
			Test("Empty json body works correctly")
		}
	})

	app.Post("/v1/post/api_json.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "{\"first\":\"foo\",\"second\":\"bar\"}" {
			Test("Setting json body using api works correctly")
		}
	})

	app.Post("/v1/post/options_json.json", func(rw http.ResponseWriter, req *http.Request) {
		if data, _ := ioutil.ReadAll(req.Body); string(data) == "{\"foo\":[\"bar\",\"baz\"]}" {
			Test("Setting json body using options works correctly")
		}
	})

	app.Patch("/v1/methods/patch.json", func() string {
		Suite("HTTP methods", 3)
		Test("Basic PATCH request is successful")
		return "/"
	})

	app.Put("/v1/methods/put.json", func() string {
		Test("Basic PUT request is successful")
		return "/"
	})

	app.Delete("/v1/methods/delete.json", func() string {
		Test("Basic DELETE request is successful")
		return "/"
	})

	go http.ListenAndServe(":3000", app)
	http.ListenAndServe(":3001", tmp)
}

func Suite(name string, count int) {
	if SuiteCount > TestsCount {
		terminal.Stdout.Color("r!").Print(" ✗ ").Color("r").Print(strconv.Itoa(SuiteCount-TestsCount) + " out of " + strconv.Itoa(SuiteCount) + " failed!").Nl()
	} else if SuiteCount < TestsCount {
		terminal.Stdout.Color("r!").Print(" ✗ ").Color("r").Print("Got " + strconv.Itoa(TestsCount-SuiteCount) + " extra tests!")
	}

	ClassCount++
	SuiteCount = count
	TestsCount = 0

	terminal.Stdout.Nl().Color("y!").Print("• ").Color("w!").Print(name).Nl()
}

func Test(name string) {
	TestsCount++
	terminal.Stdout.Color("g!").Print("\t✓ ").Color("|").Print(name).Nl()
}

func createServer() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}
