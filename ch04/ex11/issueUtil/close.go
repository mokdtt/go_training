package issueUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Close(repo, number string) (*Issue, error) {
	issue := IssueClose{"closed"}
	issueClose, err := json.Marshal(&issue)
	if err != nil {
		return nil, err
	}

	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "issues", number}, "/")

	req, err := http.NewRequest("PATCH", apiURL, bytes.NewBuffer([]byte(issueClose)))
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
		return nil, fmt.Errorf("close failed %s: %s", apiURL, resp.Status)
	}

	var issueClosed Issue
	if err := json.NewDecoder(resp.Body).Decode(&issueClosed); err != nil {
		return nil, err
	}
	return &issueClosed, nil

}
