package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repo struct {
	StarCount int `json:"stargazers_count"`
}

func GetAverageStarsPerRepo(username string) (float64, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return 0, err
	}
	// fmt.Println("RES.Body", res.Body)

	repos := []Repo{}
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
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

func hehe() {
	var client http.Client
	resp, err := client.Get("https://api.github.com/users/octocat/repos")
	if err != nil {
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}
		bodyString := string(bodyBytes)
		fmt.Println("RES.Body", bodyString)

	}
}

func main() {
	// hehe()
}
