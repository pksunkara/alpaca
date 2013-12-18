package main

import (
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

	app.Get("/v1/", func() string {
		Suite("Client Options")
		Test("Base value given by api")
		return "/"
	})

	tmp.Get("/v1/", func() string {
		Test("Base value given by options")
		return "/"
	})

	go http.ListenAndServe(":3000", app)
	http.ListenAndServe(":3001", tmp)
}
