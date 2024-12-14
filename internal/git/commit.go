package git

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	// ErrEmptyMessage is returned when commit message is empty
	ErrEmptyMessage = errors.New("commit message is empty")
	// ErrCommitNotFound is returned when commit file is not found
	ErrCommitNotFound = errors.New("commit file is not found")
)

var messagePath = filepath.Join(".git", "COMMIT_EDITMSG")

// CommitMessage represents commit message
type CommitMessage struct {
	Title string
	Body  string
}

// String returns string representation of commit message
func (m CommitMessage) String() string {
	var s strings.Builder
	s.WriteString(m.Title)
	if m.Body != "" {
		s.WriteString("\n\n")
		s.WriteString(m.Body)
	}
	return s.String()
}

// ParseCommitMessage parses commit message to parts
func ParseCommitMessage(text string) (CommitMessage, error) {
	if len(strings.Trim(text, " \n\t")) == 0 {
		return CommitMessage{}, ErrEmptyMessage
	}
	index := strings.Index(text, "\n\n")
	if index == -1 {
		return CommitMessage{
			Title: text,
		}, nil
	}
	return CommitMessage{
		Title: text[:index],
		Body:  text[index+2:],
	}, nil
}

// GetCommitMessage reads and parses commit message from internal git file.
func ReadCommitMessage() (CommitMessage, error) {
	content, err := os.ReadFile(messagePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrCommitNotFound
		}
		return CommitMessage{}, err
	}
	return ParseCommitMessage(string(content))
}

// WriteCommitMessage writes commit message to internal git file.
func WriteCommitMessage(message CommitMessage) error {
	return os.WriteFile(messagePath, []byte(message.String()), 0644)
}
