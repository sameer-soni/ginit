// Note: Only support simple relative paths like 'file.txt' — no './' or '../' or absolute paths for now.
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
		var output string = commands.HashObjectCommand(flag, file)
		fmt.Println("Hash: ", output)
	case "cat-file":
		commands.CatFileCommand(flag, file) // file = hashid here
	case "add":
		commands.AddCommand(file)
	case "status":
		commands.StatusCommand()
	default:
		fmt.Println("Unknown command: ", command)
	}
}
