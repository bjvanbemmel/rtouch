package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var force bool

func main() {
	initFlags()

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

//  Confirmation prompt
func askConsent() bool {
	var consent string

	fmt.Println("New folder(s) will have to be created. Do you consent? [y/N]")
	fmt.Scan(&consent)

	consent = strings.ToLower(consent)

	switch consent {
	case "y", "yes":
		return true
	default:
		return false
	}
}

// Separate file and folder(s) from path, execute create functions and catch errors.
func createPath(givenFilePath string) []error {
	folder, file := filepath.Split(givenFilePath)

	var allErrors = []error{}

	if _, err := os.Stat(folder); os.IsNotExist(err) && folder != "" {
		if !force {
			if !askConsent() {
				allErrors = append(allErrors, fmt.Errorf("permission not granded - canceled"))
			}
		}
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
	flag.BoolVar(&force, "f", false, "Skip confirmation prompt if folder(s) do not exist yet.")
	flag.BoolVar(&force, "force", false, "Skip confirmation prompt if folder(s) do not exist yet.")

	var helpUsage string = "Show all available commands with concise explanations."
	var forceUsage string = "Skip confirmation prompt when creating new folder(s)."

	flag.Usage = func() {
		fmt.Println("Usage: rtouch [--force] [FILEPATH]")
		fmt.Println("Resursive Touch (rtouch) allows you to create a new file and/or the corresponding folder(s) when they don't exist yet.")
		fmt.Printf("\n-h, --help %58v\n", helpUsage)
		fmt.Printf("-f, --force %56v\n\n", forceUsage)
		fmt.Println("(C) Beau Jean van Bemmel, 2022")
	}

	flag.Parse()
}
