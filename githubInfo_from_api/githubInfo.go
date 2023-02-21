package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	name, numRepos, err := githubInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println("Done\n", name, numRepos)
}

func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, err
	}
	var user gitHubUser
	userDecoder := json.NewDecoder(resp.Body)
	if err := userDecoder.Decode(&user); err != nil {
		return "", 0, err
	}
	return user.Name, user.NumRepos, nil
}

type gitHubUser struct {
	Name string
	//Public_Repos int
	NumRepos int `json:"public_repos,omitempty"` // field tag
}
