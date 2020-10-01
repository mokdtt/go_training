package githubhtml

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetContributor(repo string) (*[]Contributor, error) {
	apiURL := strings.Join([]string{restAPIURL, "repos", repo, "contributors"}, "/")
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s", apiURL, resp.Status)
	}

	var cb []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&cb); err != nil {
		return nil, err
	}
	return &cb, nil
}
