package githubhtml

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(repo string) (*IssuesSearchResult, error) {
	resp, err := http.Get(IssuesURL + "?q=repo:" + repo)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
