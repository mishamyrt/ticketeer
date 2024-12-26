package cmd

import (
	"fmt"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
)

// Ticketeer is a command line utility to add ticket id to commit message
func Ticketeer() int {
	rootCmd := newRootCmd()

	if !git.IsAvailable() {
		m := "git is not available in PATH.\n" +
			"ticketeer can't work without it.\n"
		fmt.Println(m)
		return 1
	}

	err := rootCmd.Execute()
	if err == nil {
		return 0
	}
	if !ticketeer.IsHandledError(err) {
		fmt.Printf("Unexpected error: %v\n", err)
	}
	return 1
}
