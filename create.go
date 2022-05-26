package main

import (
	"log"
	"os"
)

func createFile(filePath string) {
	_, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}
}

func createFolder(filePath string) {
	err := os.MkdirAll(filePath, 0755)

	if err != nil {
		log.Fatal(err)
	}
}
