package issueUtil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Read(repo string, number string) (*Issue, error) {
	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "issues", number}, "/")
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s", apiURL, resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func ShowIssue(issue *Issue, repo string, number string) {
	fmt.Printf("Repository: %s\nIssue Number: %s\n", repo, number)
	fmt.Printf("Title: %s\nCreated at %s\n", issue.Title, issue.CreatedAt)
	fmt.Println(issue.Body)
}
