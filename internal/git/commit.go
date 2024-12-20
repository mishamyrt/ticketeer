package git

import (
	"errors"
	"strings"
)

var (
	// ErrEmptyMessage is returned when commit message is empty
	ErrEmptyMessage = errors.New("commit message is empty")
	// ErrCommitNotFound is returned when commit file is not found
	ErrCommitNotFound = errors.New("commit file is not found")
)

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

// Bytes returns byte representation of commit message
func (m CommitMessage) Bytes() []byte {
	return []byte(m.String())
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
