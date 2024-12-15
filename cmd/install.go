package cmd

import (
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
	"github.com/spf13/cobra"
)

type install struct{}

func (install) New(opts *ticketeer.Options) *cobra.Command {
	force := false

	applyCmd := cobra.Command{
		Use:     "install",
		Short:   "Install git hook",
		Example: "ticketeer install",
		Args:    cobra.MaximumNArgs(0),
		RunE: func(_ *cobra.Command, _ []string) error {
			return ticketeer.Install(opts, force)
		},
	}

	applyCmd.Flags().BoolVarP(
		&force, "force", "f", false,
		"force installation, overwrite existing hook",
	)

	return &applyCmd
}
