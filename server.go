package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
)

func createServer() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}

func main() {
	app := createServer()
	tmp := createServer()

	app.Get("/v1/", func(rw http.ResponseWriter, req *http.Request) {
		Suite("Client Options")

		Test("Base value given by api")
		Test("Api version value given by api")

		if req.UserAgent() == "alpaca/0.1.0 (https://github.com/pksunkara/alpaca)" {
			Test("User agent value given by api")
		}

		fmt.Fprintf(rw, "/")
	})

	app.Get("/v1/method", func() string {
		Suite("Methods")
		Test("GET request")

		return "/"
	})

	app.Post("/v1/method", func() string {
		Test("POST request")
		return "/"
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

		fmt.Fprintf(rw, "/")
	})

	go http.ListenAndServe(":3000", app)
	http.ListenAndServe(":3001", tmp)
}
