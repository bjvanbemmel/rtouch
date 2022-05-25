package main

import (
	"fmt"
	"os"
)

func checkArgs() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("No filepath given.")
		os.Exit(1)
	}
}

/*
* If force flag has been given, start looking at Args position 2.
 */

func checkArgsWithFlag() {
	if forced {
		if len(os.Args[2:]) == 0 {
			fmt.Println("No filepath given.")
			os.Exit(1)
		}
	}
}

func checkFolderExists(folders string) bool {
	_, err := os.Stat(folders)

	return err == nil
}
