package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	initFlags()

	// If no arguments have been given, exit the program with status code 1.
	if len(os.Args[1:]) == 0 {
		fmt.Println("No filepath given.")
		os.Exit(1)
	}

	for _, path := range os.Args[1:] {
		if !strings.HasPrefix(path, "-") {
			for _, err := range createPath(path) {
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

// Separate file and folder(s) from path, execute create functions and catch errors.
func createPath(givenFilePath string) []error {
	folder, file := filepath.Split(givenFilePath)

	var allErrors = []error{}

	// If folders do not yet exist, create them
	if _, err := os.Stat(folder); os.IsNotExist(err) && folder != "" {
		allErrors = append(allErrors, createFolder(folder))
	}

	allErrors = append(allErrors, createFile(folder, file))

	return allErrors
}

// Create file, return error
func createFile(givenFolder, givenFile string) error {
	_, err := os.Create(givenFolder + givenFile)
	return err
}

// Create folder(s), return error
func createFolder(givenFolder string) error {
	err := os.MkdirAll(givenFolder, 0755)
	return err
}

// Create the flags and usage instructions.
func initFlags() {
	var helpUsage string = "Show all available commands with concise explanations."

	flag.Usage = func() {
		fmt.Println("Usage: rtouch [--force] [FILEPATH]")
		fmt.Println("Resursive Touch (rtouch) allows you to create a new file and/or the corresponding folder(s) when they don't exist yet.")
		fmt.Printf("\n-h, --help %58v\n", helpUsage)
		fmt.Println("(C) Beau Jean van Bemmel, 2022")
	}

	flag.Parse()
}
