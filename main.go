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

	fmt.Println(alpaca.ReadFile("config.json"))

	fmt.Println(args)
}
