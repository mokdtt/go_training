package issueUtil

import (
	"time"
)

const restAPIURL = "https://api.github.com"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssueCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type IssueClose struct {
	State string `json:"state"`
}

type IssueEdit struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
