package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dryRun = kingpin.Flag("dry-run", "Dry run").Short('n').Bool()
)

func main() {
	kingpin.Parse()

	if *dryRun {
		fmt.Printf("[Dry run] invoke process... done")
		return;
	}

	fmt.Printf("invoke process... done")
}
