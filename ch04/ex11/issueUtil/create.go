package issueUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Create(repo, title, body string) (*Issue, error) {
	issue := IssueCreate{title, body}
	issueNew, err := json.Marshal(&issue)
	if err != nil {
		return nil, err
	}

	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "issues"}, "/")

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(issueNew)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create failed %s: %s", apiURL, resp.Status)
	}

	var issueCreated Issue
	if err := json.NewDecoder(resp.Body).Decode(&issueCreated); err != nil {
		return nil, err
	}
	return &issueCreated, nil

}
