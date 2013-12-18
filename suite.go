package main

import (
	"github.com/wsxiaoys/terminal"
)

func Suite(name string) {
	terminal.Stdout.Nl().Color("y!").Print("• ").Color("w!").Print(name).Nl()
}

func Test(name string) {
	terminal.Stdout.Color("g!").Print("\t✓ ").Color("|").Print(name).Nl()
}
