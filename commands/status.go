package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func colorText(text, color string) string {
	switch color {
	case "red":
		return "\033[31m" + text + "\033[0m"
	case "green":
		return "\033[32m" + text + "\033[0m"
	case "yellow":
		return "\033[33m" + text + "\033[0m"
	case "white":
		return "\033[37m" + text + "\033[0m"
	default:
		return text
	}
}

func printSection(title, icon, symbol, color string, files []string) {
	if len(files) == 0 {
		return
	}

	header := icon + " " + title
	fmt.Println(colorText(header, color))

	for _, file := range files {
		fmt.Println(" ", symbol, " ", file)
	}
	fmt.Println()
}

func includeString(slice []string, str string) bool {
	// for _, string := range slice {
	// 	if string == str {
	// 		return true
	// 	}
	// }
	//
	// return false

	return slices.Contains(slice, str)
}

func StatusCommand() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Err getting working directory")
		return
	}

	ginitDir := filepath.Join(cwd, ".ginit")
	// checking if .ginit file exist or not
	_, err = os.Stat(ginitDir)
	if err != nil {
		fmt.Println("ginit has not been intialized")
		return
	}

	file, err := os.Open(filepath.Join(ginitDir, "index"))
	if err != nil {
		fmt.Println("Error reading index file")
		return
	}

	deletedFiles := []string{}
	filesModified := []string{}
	untrackedFiles := []string{}
	addedFiles := []string{}

	indexedFiles := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		indexedFiles = append(indexedFiles, parts[0])

		// fmt.Printf("part1-> %s, part2-> %s \n", parts[0], parts[1])

		// checking if file is deleted or not
		_, err := os.Stat(parts[0])
		if os.IsNotExist(err) {
			// fmt.Printf("File deleted: %s\n", parts[0])
			deletedFiles = append(deletedFiles, parts[0])
			continue
		}

		// checking changes in current and staging(index)
		currentHash := HashObjectCommand("", parts[0])

		if currentHash != parts[1] {
			// fmt.Println("File modified: ", parts[0])
			filesModified = append(filesModified, parts[0])
		} else {
			addedFiles = append(addedFiles, parts[0])
		}
	}

	// checking untracked files
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Err accessing path: ", err)
			return nil
		}

		if strings.Contains(path, ".ginit") {
			return nil
		}

		relPath, _ := filepath.Rel(cwd, path)
		if relPath == "." {
			return nil
		}
		tracked := includeString(indexedFiles, relPath)

		if !tracked {
			untrackedFiles = append(untrackedFiles, relPath)
		}

		// fmt.Println("path: ", relPath)
		return nil
	})
	if err != nil {
		fmt.Println("Error walking: ", cwd)
	}

	fmt.Println()
	printSection("Staged files (to be committed):", "🟢", "[+]", "green", addedFiles)
	printSection("Modified Files (not staged):", "🟡", "[~]", "yellow", filesModified)
	printSection("Deleted files:", "🔴", "[-]", "red", deletedFiles)
	printSection("Untracked files:", "⚪", "[?]", "white", untrackedFiles)
}
