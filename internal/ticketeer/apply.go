package ticketeer

import (
	"errors"
	"fmt"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/ticketeer/format"
)

// ApplyArgs represent arguments for apply command
type ApplyArgs struct {
	DryRunWith string
}

// Apply appends ticket id to commit message
func Apply(opts *Options, args *ApplyArgs) error {
	cfg, err := config.FromYAMLFile(opts.ConfigPath)
	if err != nil &&
		(!errors.Is(err, config.ErrFileNotFound) ||
			opts.ConfigPath != config.DefaultPath) {
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

	rawID, err := ticket.FindInBranch(
		branchName.String(),
		cfg.Branch.Format,
	)
	if err != nil {
		return err
	}

	id, err := ticket.ParseID(rawID, cfg.Ticket.Format)
	if err != nil {
		return err
	}

	err = format.Message(&message, id, cfg.Message)
	if err != nil {
		return err
	}

	if args.DryRunWith != "" {
		fmt.Println(message.String())
		return nil
	}

	return git.WriteCommitMessage(message)
}
