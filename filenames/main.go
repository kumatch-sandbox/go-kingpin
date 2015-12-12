package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	files = kingpin.Arg("name", "Input name").Required().ExistingFiles()
)

func main() {
	kingpin.Parse()

	for _, file := range *files {
		fmt.Printf("Exists file: %s\n", file)
	}
}
