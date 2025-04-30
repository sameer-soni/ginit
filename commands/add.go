package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddCommand(file string) {
	if file == "" {
		panic("Error: no file provided. exiting...")
	}

	// creating hash and creating file in /objects folder
	var hash string = HashObjectCommand("-w", file)

	// getting current directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting cwd: ", err)
		return
	}

	// updating index file
	f, err := os.OpenFile(filepath.Join(cwd, ".ginit", "index"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while getting index: ", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s %s\n", file, hash))
	if err != nil {
		fmt.Println("Error while writing index: ", err)
	}

	fmt.Println(file, " file has been added.", file)
}
