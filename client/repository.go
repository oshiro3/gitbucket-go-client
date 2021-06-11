package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserRepoParam struct {
	UserType  string `json:"type"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	PerPage   int    `json:"per_page"`
	Page      int    `json:"page"`
}

func (g *Gitbucket) GetPublicRepositories() ([]Repository, error) {
	res, err := http.Get(g.buildGetPublicRepositoriesUrl())
	if res.StatusCode != 200 || err != nil {
		return nil, errors.New(fmt.Sprintf("request do not succee %#v", res))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (g *Gitbucket) GetUserRepositories(user string, param *UserRepoParam) ([]Repository, error) {
	pbody, err := json.Marshal(param)
	res, err := g.requestWithoutAuth(
		"GET",
		g.buildGetUserRepositoriesUrl(user),
		strings.NewReader(string(pbody)),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("request do not succeed %#v", err))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (g *Gitbucket) buildGetUserRepositoriesUrl(user string) string {
	return fmt.Sprintf("%s/users/%s/repos", g.baseUrl(), user)
}

func (g *Gitbucket) buildGetPublicRepositoriesUrl() string {
	return fmt.Sprintf("%s/repositories", g.baseUrl())
}
