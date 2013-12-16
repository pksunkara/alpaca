package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	app := martini.Classic()

	app.Get("/", func() string {
		return "hello"
	})

	app.Run()
}
