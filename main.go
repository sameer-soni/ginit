package main

import (
	"fmt"
	"os"

	"github.com/sameer-soni/ginit/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No commands provided")
		return
	}

	if len(os.Args) > 4 {
		fmt.Println("Too many arguments provided. Exiting..")
		return
	}

	command := os.Args[1]
	var file, flag string

	if len(os.Args) == 4 {
		flag = os.Args[2]
		file = os.Args[3]
	} else if len(os.Args) == 3 {
		file = os.Args[2]
	}

	switch command {
	case "init":
		commands.InitCommand()
	case "hash-object":
		commands.HashObjectCommand(flag, file)
	default:
		fmt.Println("Unknown command: ", command)
	}
}
