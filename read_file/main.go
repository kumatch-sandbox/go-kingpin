package main

import (
	"fmt"
	"bufio"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file = kingpin.Arg("file", "Input filename").Required().File()
)

func main() {
	kingpin.Parse()

	defer func () {
		(*file).Close()
	}()

	scanner := bufio.NewScanner(*file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
