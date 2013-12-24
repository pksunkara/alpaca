package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/wsxiaoys/terminal"
	"net/http"
)

func main() {
	app := createServer()
	tmp := createServer()

	app.Get("/v1/", func(rw http.ResponseWriter, req *http.Request) {
		Suite("Client Options")
		Test("Base value given by api")
		Test("Api version value given by api")
		Test("Suffix defined by api")

		if req.UserAgent() == "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)" {
			Test("User agent value given by api")
		}

		fmt.Fprintf(rw, "/")
	})

	app.Get("/v1/method", func() string {
		Suite("HTTP Methods")
		Test("GET request")
		return "/"
	})

	app.Post("/v1/method", func() string {
		Test("POST request")
		return "/"
	})

	app.Put("/v1/method", func() string {
		Test("PUT request")
		return "/"
	})

	app.Patch("/v1/method", func() string {
		Test("PATCH request")
		return "/"
	})

	app.Delete("/v1/method", func() string {
		Test("DELETE request")
		return "/"
	})

	app.Get("/v2/", func(rw http.ResponseWriter, req *http.Request) {
		Suite("Request Options")
		Test("Base value should not be changed")
		Test("Api version value given by options")

		if req.UserAgent() == "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)" {
			Test("User agent value should not be changed")
		}

		if req.Header["Custom-Header"][0] == "custom" {
			Test("Headers value given by options")
		}

		fmt.Fprintf(rw, "/")
	})

	tmp.Get("/v2/", func(rw http.ResponseWriter, req *http.Request) {
		Test("Base value given by options")
		Test("Api version value given by options")

		if req.UserAgent() == "testing (user agent)" {
			Test("User agent value given by options")
		}

		if req.Header["Custom-Header"][0] == "custom" {
			Test("Headers value given by options")
		}

		Test("Base value's path should not be used")

		fmt.Fprintf(rw, "/")
	})

	go http.ListenAndServe(":3000", app)
	http.ListenAndServe(":3001", tmp)
}

func Suite(name string) {
	terminal.Stdout.Nl().Color("y!").Print("• ").Color("w!").Print(name).Nl()
}

func Test(name string) {
	terminal.Stdout.Color("g!").Print("\t✓ ").Color("|").Print(name).Nl()
}

func createServer() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}
