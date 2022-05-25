package main

import (
	"flag"
	"fmt"
)

var forced bool

func main() {
	initFlags()

	if forced {
		fmt.Println("Forced")
	}
}

func initFlags() {
	flag.BoolVar(&forced, "f", false, "Skip confirmation prompt if folder(s) do not yet exist")
	flag.BoolVar(&forced, "force", false, "Skip confirmation prompt if folder(s) do not yet exist")

	initUsageHelp()

	flag.Parse()
}

func initUsageHelp() {
	var helpUsage string = "Show all available commands with concise explanations."
	var forceUsage string = "Skip confirmation prompt to create new folders when folder(s) do not exist yet."

	flag.Usage = func() {
		fmt.Printf("Usage: rtouch [--force] [FILEPATH]\n")
		fmt.Printf("Recursive Touch (rtouch) allows you to create a new file and/or the corresponding folder(s) when it doesn't exist yet.\n\n")
		fmt.Printf("-h, --help %58v\n", helpUsage)
		fmt.Printf("-f, --force %82v\n\n", forceUsage)
		fmt.Println("(C) Beau Jean van Bemmel, 2022")
	}
}
