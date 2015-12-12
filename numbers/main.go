package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	numbers = kingpin.Arg("numbers", "some numbers").Uint8List()
)

func main() {
	kingpin.Parse()
	fmt.Printf("numbers: %v", *numbers)
}
