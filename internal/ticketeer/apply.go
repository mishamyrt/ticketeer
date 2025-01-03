package ticketeer

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/ticketeer/format"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
	"github.com/mishamyrt/ticketeer/pkg/pattern"
)

// ApplyArgs represent arguments for apply command
type ApplyArgs struct {
	ConfigPath string
	DryRunWith string
}

// Apply appends ticket id to commit message
func (a *App) Apply(workingDir string, args *ApplyArgs) error {
	cfg, err := a.resolveConfig(workingDir, args.ConfigPath)
	if err != nil {
		return a.handleError(err, "Failed to resolve config")
	}

	repo, err := git.OpenRepository(workingDir)
	if err != nil {
		return a.handleError(err, "Failed to open repository")
	}
	a.log.Debugf("Repository root found at: %s", repo.Path())

	branchName, err := repo.BranchName()
	if err != nil {
		a.log.Info("Branch is not found, skipping")
		a.log.Debugf("Error: %v", err)
		return nil
	}
	a.log.Debugf("Branch name: %s", branchName)

	ignores := pattern.NewList(cfg.Branch.Ignore...)
	if ignores.Match(branchName) {
		a.log.Info("Branch is ignored, skipping")
		return nil
	}

	var message git.CommitMessage
	if args.DryRunWith != "" {
		message, err = git.ParseCommitMessage(args.DryRunWith)
	} else {
		message, err = repo.CommitMessage()
	}
	if err != nil {
		return a.handleError(err, "Failed to parse commit message")
	}

	rawID, err := ticket.FindInBranch(branchName, cfg.Branch.Format)
	if err != nil {
		return a.handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}
	id, err := ticket.ParseID(rawID, cfg.Ticket.Format)
	if err != nil {
		return a.handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}
	a.log.Debugf("Ticket ID found in branch name: %s", rawID)

	err = format.Message(&message, id, cfg.Message)
	if err != nil {
		return a.handleError(err, "Failed to format commit message")
	}

	if args.DryRunWith != "" {
		a.log.Info("Running in dry-run mode")
		a.log.Info(message.String())
		return nil
	}

	err = repo.SetCommitMessage(message)
	if err != nil {
		return a.handleError(err, "Failed to update commit message")
	}

	return nil
}

func (a *App) resolveConfig(workingDir, path string) (*config.Config, error) {
	var configPath string
	if path == "" {
		configPath = filepath.Join(workingDir, config.DefaultFileName)
	} else if filepath.IsAbs(path) {
		configPath = path
	} else {
		configPath = filepath.Join(workingDir, path)
	}

	cfg, err := config.FromYAMLFile(configPath)
	if err == nil {
		return cfg, nil
	}

	if errors.Is(err, config.ErrUnknownLocation) {
		a.log.Info(color.Red("Unknown ticket location"))
		a.log.Info("Available options:")
		for _, v := range config.TicketLocationOptions {
			a.log.Info(fmt.Sprintf("- %s", v))
		}
	}

	if errors.Is(err, config.ErrUnknownBranchFormat) {
		a.log.Info(color.Red("Unknown branch format"))
		a.log.Info("Available options:")
		for _, v := range config.BranchFormatOptions {
			a.log.Info(fmt.Sprintf("- %s", v))
		}
	}

	if errors.Is(err, config.ErrUnknownTicketFormat) {
		a.log.Info(color.Red("Unknown ticket format"))
		a.log.Info("Available options:")
		for _, v := range ticket.IDFormatOptions {
			a.log.Info(fmt.Sprintf("- %s", v))
		}
	}

	if errors.Is(err, config.ErrFileNotFound) && path == "" {
		return &config.Default, nil
	}
	return nil, err
}

func (a *App) handleEmptyTicket(err error, allowEmpty bool) error {
	if !allowEmpty {
		return a.handleError(err, "Ticket ID is not found in branch name")
	}
	a.log.Info("Ticket ID is not found in branch name, skipping")
	return nil
}
