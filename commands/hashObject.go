package commands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func HashObjectCommand(flag string, file string) string {
	// fmt.Println("in hash object")

	if flag != "" && flag != "-w" {
		fmt.Println("invalid flag.")
		fmt.Println("flag available:  -w  -> write file")
		return ""
	}

	if file == "" {
		fmt.Println("file name missing. Exiting...")
		return ""
	}

	cwd, error := os.Getwd()
	if error != nil {
		fmt.Println("Error getting current directory: ", error)
		return ""
	}

	data, err := os.ReadFile(filepath.Join(cwd, file))
	if err != nil {
		fmt.Println("Error while reading file: ", err)
		return ""
	}

	textContent := string(data)

	// hashing - SHA1
	header := fmt.Sprintf("blob %d", len(textContent))
	blob := append([]byte(header), 0)
	blob = append(blob, []byte(textContent)...)

	hasher := sha1.New()
	hasher.Write(blob)
	hashedBytes := hasher.Sum(nil)

	hexValue := fmt.Sprintf("%x", hashedBytes)

	// compression- zlib
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	// _, errCompression := w.Write([]byte(textContent))
	_, errCompression := w.Write(blob)

	if errCompression != nil {
		fmt.Println("Error during compression: ", err)
		return ""
	}
	w.Close()

	// saving it in objects folder

	if flag == "-w" {
		// check if ginit is initialized or not
		_, gerr := os.Stat(filepath.Join(cwd, ".ginit"))
		if os.IsNotExist(gerr) {
			fmt.Println("ginit has not been initialized")
			return ""
		} else if gerr != nil {
			fmt.Println("Error gerr: ", gerr)
			return ""
		}

		subFolder := hexValue[:2]

		err = os.MkdirAll(filepath.Join(cwd, ".ginit", "objects", subFolder), 0755)
		if err != nil {
			fmt.Println("Error while creating subfolder: ", err)
			return ""
		}

		f, err := os.Create(filepath.Join(cwd, ".ginit", "objects", subFolder, hexValue[2:]))
		if err != nil {
			fmt.Println("Err while creating file: ", err)
			return ""
		}
		defer f.Close()
		_, err = f.Write(b.Bytes())
		if err != nil {
			fmt.Println("Err while writing compressed blob: ", err)
		}

		// fmt.Println("object file created: ", hexValue)

	}

	// fmt.Println("Object hash: ", hexValue)
	return hexValue
}
