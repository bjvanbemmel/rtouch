package main

import (
	"os"
)

func initFilePath() string {
	if forced {
		return os.Args[2]
	} else {
		return os.Args[1]
	}

}
