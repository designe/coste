package main

import (
	"fmt"
	"os"
)

var commands = make(map[string]func())

/*
Initialize Reserved Commands
*/
func makeReservedCommand() {
	Copy := func() {
		fmt.Println("Call Copy!")
	}

	Paste := func() {
		fmt.Println("Call Paste!")
	}

	Help := func() {
		fmt.Println("Help!")
	}
	commands["c"] = Copy
	commands["-c"] = Copy
	commands["copy"] = Copy
	commands["--copy"] = Copy

	commands["v"] = Paste
	commands["-v"] = Paste
	commands["paste"] = Paste
	commands["--paste"] = Paste

	commands["help"] = Help
}

func main() {
	// CHECK Environment Variables : COSTE_HOME
	if coste_home := os.Getenv("COSTE_HOME"); coste_home == "" {
		fmt.Println("Environment Variable \"COSTE_HOME\" is not defined.")
		return
	} else {
		fmt.Println("COSTE_HOME = " coste_home)
	}
	
	makeReservedCommand()
	
	args := os.Args[1:]

	command := args[0]

	commandFunc, presence := commands[command]

	if presence {
		fmt.Println(command + "is present.")
		commandFunc()
		
	} else {
		fmt.Println(command + "is not present.")
	}
}
