package main

import (
	"fmt"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("manager", "Management command")

	itemCommand = app.Command("item", "Item management")
	// item add
	itemAdd = itemCommand.Command("add", "Add item")
	itemAddName = itemAdd.Flag("name", "item name").Required().String()
	// item del
	itemDel = itemCommand.Command("del", "Delete item")
	itemDelName  = itemDel.Flag("name", "item name").Required().String()
	itemDelForce = itemDel.Flag("force", "force delete").Short('f').Bool()

	personCommand = app.Command("person", "Person management")
	// item add
	personAdd = personCommand.Command("add", "Add person")
	personAddName = personAdd.Flag("name", "person name").Required().String()
)

func main() {

	//switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case itemAdd.FullCommand():
		fmt.Printf("[Add item] name: %s\n", *itemAddName)
	case itemDel.FullCommand():
		forceDelete := ""
		if *itemDelForce {
			forceDelete = " (force)"
		}

		fmt.Printf("[Delete item%s] name: %s\n", forceDelete, *itemDelName)
	case personAdd.FullCommand():
		fmt.Printf("[Add person] name: %s\n", *personAddName)
	}
}
