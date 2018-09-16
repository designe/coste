package main

import (
	"fmt"
	"os"
)

var commands = make(map[string]bool)

func makeReservedCommand() {
	commands["c"] = true
	commands["-c"] = true
	commands["copy"] = true
	commands["--copy"] = true

	commands["v"] = true
	commands["-v"] = true
	commands["paste"] = true
	commands["--paste"] = true

	commands["help"] = true
}

func main() {
	makeReservedCommand()
	
	args := os.Args[1:]

	command := args[0]

	_, presence := commands[command]

	if presence == true {
		fmt.Println(command + "is present.")
	} else {
		fmt.Println(command + "is not present.")
	}
}
