package simple

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

func GetAverageStarsPerRepo(username string) (float64, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return 0, err
	}

	repos := []Repo{}
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return 0, err
	}

	if len(repos) == 0 {
		return 0, nil
	}

	var total int
	for _, r := range repos {
		total += r.StargazersCount
	}

	return float64(total) / float64(len(repos)), nil
}

func TestGetAverageStarsPerRepo(t *testing.T) {
	var tests = []struct {
		username string
		want     float64
	}{
		{"octocat", 1480.375000},
		{"plutov", 15.566667},
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			got, err := GetAverageStarsPerRepo(tt.username)
			// Don't omit errors even in tests
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}
			if got != tt.want {
				t.Errorf("expecting %f, got %f", tt.want, got)
			}
		})
	}
}
