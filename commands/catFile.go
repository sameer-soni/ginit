package commands

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
)

func CatFileCommand(flag string, hash string) {
	if flag == "" {
		fmt.Println("No flag provided. Available flags: ")
		fmt.Println("-p  pretty-print the content")
		return
	}

	// valid flags:
	validFlags := []string{"-p"}

	// checking for valid flags:
	if !slices.Contains(validFlags, flag) {
		fmt.Println("Invalid flag: ", flag)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Err getting current directory")
		return
	}

	objectPath := filepath.Join(cwd, ".ginit", "objects")

	fileFullPath := filepath.Join(objectPath, hash[0:2], hash[2:])

	data, err := os.ReadFile(fileFullPath)
	if err != nil {
		fmt.Println("Error reading hash file: ", err)
		return
	}

	// decompressing
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		fmt.Println("error while decompression: ", err)
		return

	}

	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil {
		fmt.Println("Error while copying buffer: ", err)
		return
	}

	if flag == "-p" {
		fmt.Println(out.String())
	}
}
