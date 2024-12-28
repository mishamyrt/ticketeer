package cmd

import (
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type version struct{}

func (version) New(app *ticketeer.App) *cobra.Command {
	var full bool

	versionCmd := cobra.Command{
		Use:     "version",
		Short:   "Print the version",
		Example: "ticketeer version",
		Args:    cobra.MaximumNArgs(0),
		RunE: func(_ *cobra.Command, _ []string) error {
			app.Version(full)
			return nil
		},
	}

	versionCmd.Flags().BoolVarP(
		&full, "full", "f", false,
		"print full information (commit hash)",
	)

	return &versionCmd
}
