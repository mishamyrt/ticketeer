package cmd

import (
	"os"

	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type apply struct{}

func (apply) New(app *ticketeer.App) *cobra.Command {
	var applyArgs ticketeer.ApplyArgs

	applyCmd := cobra.Command{
		Use:     "apply",
		Short:   "Append ticket id to commit message",
		Example: "ticketeer apply",
		RunE: func(_ *cobra.Command, _ []string) error {
			cwd, _ := os.Getwd()
			return app.Apply(cwd, &applyArgs)
		},
	}

	applyCmd.Flags().StringVar(
		&applyArgs.DryRunWith, "dry-run-with", "",
		"skip reading and writing message, use fake commit message",
	)

	applyCmd.Flags().StringVarP(
		&applyArgs.ConfigPath, "config", "c", "",
		"path to configuration file",
	)

	return &applyCmd
}
