package main

import (
	"strings"
)

var forced bool
var consented string

func main() {
	initFlags()

	checkArgs()
	checkArgsWithFlag()

	var givenFilePath = initFilePath()
	var filePathSplit = strings.Split(givenFilePath, "/")
	var newFile = filePathSplit[len(filePathSplit)-1]
	var newFolders = givenFilePath[0 : len(givenFilePath)-len(newFile)]

	if checkFolderArgGiven(filePathSplit) {
		if !checkFolderExists(newFolders) {
			if !forced {
				checkConsentAndExitIfFalse()
			}

			createFolder(newFolders)
		}
	}

	createFile(newFolders + newFile)
}