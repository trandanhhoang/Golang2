package hehe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Repo struct {
	StarCount int `json:"stargazers_count`
}

type Mock struct{}

type Github struct{}

type RepositoryAPI interface {
	getRepos(string) ([]Repo, error)
}

func (m *Mock) getRepos(username string) ([]Repo, error) {
	// log.Fatal("hehe")

	return []Repo{
		{
			StarCount: 2,
		},
		{
			StarCount: 9,
		},
	}, nil
}

func (g *Github) getRepos(username string) ([]Repo, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}

	repos := []Repo{}
	log.Fatal("hehe")
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func GetAverageStarsPerRepo(repositoryAPI RepositoryAPI, username string) (float64, error) {
	repos, err := repositoryAPI.getRepos(username)
	if err != nil {
		return 0, err
	}

	if len(repos) == 0 {
		return 0, nil
	}

	var total int
	for _, r := range repos {
		total += r.StarCount
	}

	return float64(total) / float64(len(repos)), nil
}
