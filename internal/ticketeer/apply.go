package ticketeer

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/ticketeer/format"
	"github.com/mishamyrt/ticketeer/pkg/log"
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
		return err
	}

	repo, err := git.OpenRepository(workingDir)
	if err != nil {
		return err
	}
	log.Debugf("Repository root found at: %s", repo.Path())

	branchName, err := repo.BranchName()
	if err != nil {
		log.Info("Branch is not found, skipping")
		return nil
	}
	log.Debugf("Branch name: %s", branchName)

	matcher := branchMatcher(cfg.Branch.Ignore)
	isIgnored, err := matcher.Match(branchName)
	if err != nil {
		return err
	}

	if isIgnored {
		log.Info("Branch is ignored, skipping")
		return nil
	}

	var message git.CommitMessage
	if args.DryRunWith != "" {
		message, err = git.ParseCommitMessage(args.DryRunWith)
	} else {
		message, err = repo.CommitMessage()
	}
	if err != nil {
		return err
	}

	rawID, err := ticket.FindInBranch(branchName, cfg.Branch.Format)
	if err != nil {
		return a.handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}

	id, err := ticket.ParseID(rawID, cfg.Ticket.Format)
	if err != nil {
		return a.handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}

	log.Debugf("Ticket ID found in branch name: %s", rawID)

	err = format.Message(&message, id, cfg.Message)
	if err != nil {
		return err
	}

	if args.DryRunWith != "" {
		log.Info("Running in dry-run mode")
		print(message.String())
		return nil
	}

	return repo.SetCommitMessage(message)
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
	if err != nil {
		return nil, err
	}
	if errors.Is(err, config.ErrFileNotFound) && path == "" {
		return &config.Default, nil
	}

	return cfg, nil
}

func (a *App) handleEmptyTicket(err error, allowEmpty bool) error {
	if !allowEmpty {
		return err
	}
	log.Info("Ticket ID is not found in branch name, skipping")
	return nil
}

type branchMatcher []string

func (m branchMatcher) Match(branchName string) (bool, error) {
	for _, assertion := range m {
		if !strings.Contains(assertion, "*") {
			if assertion == branchName {
				return true, nil
			}
			continue
		}
		re, err := m.build(assertion)
		if err != nil {
			return false, err
		}
		if re.MatchString(branchName) {
			return true, nil
		}
	}
	return false, nil
}

func (m branchMatcher) build(branchPath string) (*regexp.Regexp, error) {
	forbidden := regexp.MustCompile(`[\?\.\[\]\(\)\$\^]+`)
	reText := forbidden.ReplaceAllStringFunc(branchPath, func(match string) string {
		return fmt.Sprintf("\\%s", match)
	})
	reText = strings.ReplaceAll(reText, "*", ".*")
	reText = fmt.Sprintf("^%s$", reText)
	return regexp.Compile(reText)
}
