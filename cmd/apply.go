package cmd

import (
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type apply struct{}

func (apply) New(opts *ticketeer.Options) *cobra.Command {
	var applyArgs ticketeer.ApplyArgs

	applyCmd := cobra.Command{
		Use:     "apply",
		Short:   "Append ticket id to commit message",
		Example: "ticketeer apply",
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ticketeer.Apply(opts, &applyArgs)
		},
	}

	applyCmd.Flags().StringVar(
		&applyArgs.DryRunWith, "dry-run-with", "",
		"skip reading and writing message, use fake commit message",
	)

	return &applyCmd
}
