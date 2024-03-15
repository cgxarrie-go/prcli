package models

import (
	"fmt"
	"time"
)

// PullRequest is the abstraction of a Pull Request from any provider
type PullRequest struct {
	Origin       string
	Organization string
	ID           string
	Title        string
	Description  string
	Repository   Hierarchy
	Project      Hierarchy
	Status       string
	MergeStatus  string
	CreatedBy    string
	URL          string
	IsDraft      bool
	Created      time.Time
	Link         string
}

// Hierarchy of a PR
type Hierarchy struct {
	ID   string
	Name string
	URL  string
}

// ShortenedTitle returns title shortened to maxlength...
func (p PullRequest) ShortenedTitle(maxLength int) string {

	pritntable := p.Title

	if len(pritntable) <= maxLength {
		return pritntable
	}

	shortenLenght := maxLength - 3

	title := fmt.Sprintf("%s...", pritntable[0:shortenLenght])
	return title
}
