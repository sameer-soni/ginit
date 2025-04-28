package commands

import (
	"fmt"
	"os"
)

func InitCommand() {
	// fmt.Println("init command executed")

	cwd, error := os.Getwd()
	if error != nil {
		fmt.Println("Error getting current directory: ", error)
		return
	}

	// checking if .ginit already exist or not
	_, err := os.Stat(cwd + "/.ginit")
	if err == nil {
		// path exist
		fmt.Println("ginit is already initialized")
		return
	}

	if os.IsNotExist(err) {
		// dir doesn't exist

		paths := []string{
			cwd + "/.ginit",
			cwd + "/.ginit/objects/",
			cwd + "/.ginit/refs/heads/",
		}

		for _, path := range paths {
			err = os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}

		file, err := os.Create(cwd + "/.ginit/HEAD")
		if err != nil {
			fmt.Println("Error creating HEAD file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString("ref: refs/heads/master")
		if err != nil {
			fmt.Println("Error writing in HEAD file:", err)
			return
		}

		fmt.Println(".ginit has been initialized in:", cwd)
		return
	}

	fmt.Println("Error checking directory:", err)
}
