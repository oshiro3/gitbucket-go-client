package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func (g *Gitbucket) buildGetPublicRepositoriesUrl() string {
	return fmt.Sprintf("http://%s/api/v3/repositories", g.host)
}
