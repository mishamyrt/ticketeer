package cmd

import (
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/pkg/log"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

// Ticketeer is a command line utility to add ticket id to commit message
func Ticketeer() int {
	rootCmd := newRootCmd()

	if !git.IsAvailable() {
		message := color.Yellow(
			"Warning: git is not available in PATH.\n" +
				"Ticketeer may not work correctly.\n\n")
		log.Info(message)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Errorf("Error: %v", err)
		return 1
	}

	return 0
}
