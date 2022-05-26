package main

import (
	"fmt"
	"os"
	"strings"
)

func askConsent() bool {
	fmt.Println("New folder(s) will have to be created. Do you consent? [y/N]")
	fmt.Scan(&consented)

	switch strings.ToLower(consented) {
	case "y":
		fallthrough
	case "yes":
		return true
	default:
		return false
	}
}

func checkConsentAndExitIfFalse() {
	if !askConsent() {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}
}