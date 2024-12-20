package git

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrHeadNotFound is returned when HEAD file is not found
var ErrHeadNotFound = errors.New("HEAD file is not found")

// Repository represents git repository
type Repository struct {
	root string
}

// Path returns path to the git repository
func (r *Repository) Path() string {
	return r.root
}

// HooksDir returns path to the repo hooks directory
func (r *Repository) HooksDir() (string, error) {
	cmd := Command("rev-parse", "--git-path", "hooks")
	return cmd.ExecuteAt(r.root)
}

// BranchName returns current branch name.
// If repository is in detached state, returns error
func (r *Repository) BranchName() (string, error) {
	headPath := filepath.Join(r.root, ".git", "HEAD")
	content, err := os.ReadFile(headPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrHeadNotFound
		}
		return "", err
	}

	return BranchNameFromHead(string(content))
}

// CommitMessage reads current commit message and parses it to title and body.
func (r *Repository) CommitMessage() (CommitMessage, error) {
	content, err := os.ReadFile(r.commitMessagePath())
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrCommitNotFound
		}
		return CommitMessage{}, err
	}
	return ParseCommitMessage(string(content))
}

// SetCommitMessage writes commit message to the repository.
func (r *Repository) SetCommitMessage(m CommitMessage) error {
	msgPath := r.commitMessagePath()
	content := m.Bytes()
	return os.WriteFile(msgPath, content, 0644)
}

func (r *Repository) commitMessagePath() string {
	return filepath.Join(r.root, ".git", "COMMIT_EDITMSG")
}

// NewRepository returns a new Repository instance
func NewRepository(root string) *Repository {
	return &Repository{root: root}
}

// FindRoot returns root of git repository
func FindRoot(path string) (string, error) {
	cmd := Command("rev-parse", "--show-toplevel")
	return cmd.ExecuteAt(path)
}

// OpenRepository returns a new Repository instance with root of git repository.
func OpenRepository(path string) (*Repository, error) {
	root, err := FindRoot(path)
	if err != nil {
		return nil, err
	}
	return NewRepository(root), nil
}
