package main

import (
	"fmt"
	"os"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
//	"gopkg.in/src-d/go-git.v4/plumbing"
//	"gopkg.in/src-d/go-git.v4/plumbing/object"

)

var COSTE_HOME string = ""
var commands = make(map[string]func())

/*
Initialize Reserved Commands
*/
func makeReservedCommand() {
	Copy := func() {
		fmt.Println("Call Copy on " + COSTE_HOME)
		repo, err := git.PlainOpen(COSTE_HOME)
		CheckIfError(err)

		worktree, err := repo.Worktree()
		CheckIfError(err)

		err = worktree.Checkout(&git.CheckoutOptions {
			Branch: "refs/heads/coste-history",
		})
		CheckIfError(err)

		/*
1) FILE COPY TO HISTORY FOLDER
2) GIT ADD FILE NAME
3) GIT PUSH FILE NAME
4) FINISH
*/
		
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
		fmt.Println("COSTE_HOME = " + coste_home)
		COSTE_HOME = coste_home
		/*
		repo, err := git.PlainOpen(coste_home)
		CheckIfError(err)

		Info("git rev-list HEAD --count")

		branch, err := repo.Branch("test")
		CheckIfError(err)
		fmt.Println(branch.Name)
		
		branches, err := repo.Branches()
		CheckIfError(err)

		//fmt.Println(b.Name + " " + b.Remote)
		err = branches.ForEach(func(r *plumbing.Reference) error {
			fmt.Println(r.Name())
			fmt.Println(r.Strings())

			return nil
		})

		CheckIfError(err)
		*/
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
