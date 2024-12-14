package ticketeer

import (
	"fmt"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/ticketeer/render"
)

// ApplyArgs represent arguments for apply command
type ApplyArgs struct {
	DryRunWith string
}

// Apply appends ticket id to commit message
func Apply(opts *Options, args *ApplyArgs) error {
	cfg, err := config.FromYAML(opts.ConfigPath)
	if err != nil {
		return err
	}

	branchName, err := git.ReadBranchName()
	if err != nil {
		return err
	}

	var message git.CommitMessage
	if args.DryRunWith != "" {
		message, err = git.ParseCommitMessage(args.DryRunWith)
	} else {
		message, err = git.ReadCommitMessage()
	}
	if err != nil {
		return err
	}

	ticketID, err := ticket.ParseFromBranchName(
		branchName.String(),
		ticket.AlphanumericFormat, // TODO: get format from config
	)
	if err != nil {
		return err
	}

	switch cfg.TicketLocation {
	case config.TicketLocationTitle:
		message.Title, err = render.Title(cfg.Template, message.Title, ticketID)
	case config.TicketLocationBody:
		message.Body, err = render.Body(cfg.Template, message.Body, ticketID)
	}
	if err != nil {
		return err
	}

	if args.DryRunWith != "" {
		fmt.Println(message.String())
		return nil
	}

	return git.WriteCommitMessage(message)
}
