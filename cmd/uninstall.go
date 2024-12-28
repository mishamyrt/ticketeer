package cmd

import (
	"os"

	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type uninstall struct{}

func (uninstall) New(app *ticketeer.App) *cobra.Command {
	force := false

	applyCmd := cobra.Command{
		Use:     "uninstall",
		Short:   "uninstall git hook",
		Example: "ticketeer uninstall",
		Args:    cobra.MaximumNArgs(0),
		RunE: func(_ *cobra.Command, _ []string) error {
			cwd, _ := os.Getwd()
			return app.Uninstall(cwd, force)
		},
	}

	applyCmd.Flags().BoolVarP(
		&force, "force", "f", false,
		"force uninstallation, even if another hook is installed",
	)

	return &applyCmd
}
