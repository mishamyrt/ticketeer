package cmd

import "log"

func Ticketeer() int {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		if err.Error() != "" {
			log.Printf("Error: %s", err)
		}
		return 1
	}

	return 0
}
