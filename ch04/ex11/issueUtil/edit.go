package issueUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Edit(repo, number, title, body string) (*Issue, error) {
	issue := IssueEdit{title, body}
	issueEdit, err := json.Marshal(&issue)
	if err != nil {
		return nil, err
	}

	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "issues", number}, "/")

	req, err := http.NewRequest("PATCH", apiURL, bytes.NewBuffer([]byte(issueEdit)))
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("edit failed %s: %s", apiURL, resp.Status)
	}

	var issueEdited Issue
	if err := json.NewDecoder(resp.Body).Decode(&issueEdited); err != nil {
		return nil, err
	}
	return &issueEdited, nil

}
