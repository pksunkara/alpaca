package main

import (
	"./alpaca"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	Version = "0.1.0"
)

func main() {

	// Options for flags package
	var opts struct {
		Version bool `short:"v" long:"version" description:"Show version information"`
	}

	// Parse the arguments
	args, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(0)
	}

	// Print version and exit
	if opts.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	directory := args[0]

	var module, api map[string]interface{}

	alpaca.ReadJSON(directory+"/module.json", &module)
	alpaca.ReadJSON(directory+"/api.json", &api)

	data := alpaca.Data{api, module}

	alpaca.RunTemplate("alpaca/templates/node/package.json", &data)
}
