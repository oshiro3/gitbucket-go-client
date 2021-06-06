package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Status struct {
	State       string `json:"state"`
	TargetUrl   string `json:"taget_url"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

func (g *Gitbucket) SetStatus(hash string, status *Status) (*http.Response, error) {
	body, err := json.Marshal(status)
	if err != nil {
		log.Println("fail to marshal payload")
		return nil, err
	}
	res, err := g.request(
		"POST",
		g.buildStatusURL(hash),
		strings.NewReader(string(body)),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *Gitbucket) buildStatusURL(commit string) string {
	return fmt.Sprintf("%s/repos/%s/%s/statuses/%s", g.baseUrl(), g.owner, g.repositoryName, commit)
}
