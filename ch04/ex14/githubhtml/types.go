package githubhtml

import (
	"time"
)

const restAPIURL = "https://api.github.com"
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type Milestone struct {
	Number      int
	HTMLURL     string `json:"html_url"`
	State       string
	Title       string
	Description string
	Creator     *User
	CreatedAt   time.Time `json:"created_at"`
}

type Contributor struct {
	ID            int
	HTMLURL       string `json:"html_url"`
	Contributions int
	Login         string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
