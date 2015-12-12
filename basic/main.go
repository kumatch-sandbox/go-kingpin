package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Set verbose mode").Bool()
	count   = kingpin.Flag("count", "counter").Int()
	name    = kingpin.Arg("name", "Input name").Required().String()
)

func main() {
	kingpin.Parse()
	fmt.Printf("verbose mode: %v, count: %d, name: %s\n", *verbose, *count, *name)
}
