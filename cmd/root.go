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
}

func newRootCmd() *cobra.Command {
	var options ticketeer.Options
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
			app.Setup(&options)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(
		&options.Verbose, "verbose", "v", false, "verbose output",
	)

	for _, command := range commands {
		rootCmd.AddCommand(command.New(app))
	}

	return rootCmd
}
