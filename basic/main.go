package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Set verbose mode").Short('v').Bool()
	name    = kingpin.Arg("name", "Input name").Required().String()
	numbers = kingpin.Arg("numbers", "some numbers").Uint8List()
)

func main() {
	kingpin.Parse()
	fmt.Printf("name: %s, numbers: %v, verbose mode: %v\n", *name, *numbers, *verbose)
}
