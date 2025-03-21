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
	rootDir           string
	hooksDir          string
	commitMessagePath string
}

// Path returns path to the git repository
func (r *Repository) Path() string {
	return r.rootDir
}

// HooksDir returns path to the repo hooks directory
func (r *Repository) HooksDir() string {
	return r.hooksDir
}

// Exec executes git command
func (r *Repository) Exec(cmd *Cmd) (string, error) {
	return cmd.ExecuteAt(r.rootDir)
}

// BranchName returns current branch name.
// If repository is in detached state, returns error
func (r *Repository) BranchName() (string, error) {
	headPath := filepath.Join(r.rootDir, ".git", "HEAD")
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
	content, err := os.ReadFile(r.commitMessagePath)
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
	msgPath := r.commitMessagePath
	content := m.Bytes()
	return os.WriteFile(msgPath, content, 0644)
}

// FindRoot returns root of git repository
func FindRoot(path string) (string, error) {
	cmd := Command("rev-parse", "--show-toplevel")
	return cmd.ExecuteAt(path)
}

// NewRepository creates and returns a new Repository instance
func NewRepository(root string) (*Repository, error) {
	_, err := Command("init").ExecuteAt(root)
	if err != nil {
		return nil, err
	}
	return OpenRepository(root)
}

// OpenRepository returns a new Repository instance with root of git repository.
func OpenRepository(path string) (*Repository, error) {
	rootDir, err := FindRoot(path)
	if err != nil {
		return nil, err
	}
	hooksDir, err := Command("rev-parse", "--git-path", "hooks").ExecuteAt(rootDir)
	if err != nil {
		return nil, err
	}
	return &Repository{
		rootDir:           rootDir,
		hooksDir:          filepath.Join(rootDir, hooksDir),
		commitMessagePath: filepath.Join(rootDir, ".git", "COMMIT_EDITMSG"),
	}, nil
}
