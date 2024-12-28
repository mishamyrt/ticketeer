package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type command interface {
	New(*ticketeer.App) *cobra.Command
}

var commands = [...]command{
	apply{},
	install{},
	version{},
}

func newRootCmd() *cobra.Command {
	var logOpts ticketeer.LogOptions

	app := ticketeer.New()

	rootCmd := &cobra.Command{
		Use:   "ticketeer",
		Short: "Utility to add ticket id to commit message",
		Long: heredoc.Doc(`
				After installation go to your project directory
				and execute the following command:
				ticketeer install
		`),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			app.SetupLog(logOpts)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(
		&logOpts.Verbose, "verbose", "v", false, "verbose output",
	)

	rootCmd.PersistentFlags().BoolVar(
		&logOpts.NoColor, "no-color", false, "disable color output",
	)

	for _, command := range commands {
		rootCmd.AddCommand(command.New(app))
	}

	return rootCmd
}
