package githubhtml

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetMilestone(repo string) (*[]Milestone, error) {
	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "milestones"}, "/")
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s", apiURL, resp.Status)
	}

	var ms []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&ms); err != nil {
		return nil, err
	}
	return &ms, nil
}
