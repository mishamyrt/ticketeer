package ticketeer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

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

	repo, err := git.OpenRepository(".")
	if err != nil {
		return err
	}

	branchName, err := repo.BranchName()
	if err != nil {
		fmt.Println("Branch is not found, skipping")
		return nil
	}

	matcher := branchMatcher(cfg.Branch.Ignore)
	isIgnored, err := matcher.Match(branchName)
	if err != nil {
		return err
	}

	if isIgnored {
		fmt.Println("Branch is ignored, skipping")
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
		return handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}

	id, err := ticket.ParseID(rawID, cfg.Ticket.Format)
	if err != nil {
		return handleEmptyTicket(err, cfg.Ticket.AllowEmpty)
	}

	err = format.Message(&message, id, cfg.Message)
	if err != nil {
		return err
	}

	if args.DryRunWith != "" {
		fmt.Println(message.String())
		return nil
	}

	return repo.SetCommitMessage(message)
}

func handleEmptyTicket(err error, allowEmpty bool) error {
	if !allowEmpty {
		return err
	}
	fmt.Println("Ticket ID is not found in branch name, skipping")
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
