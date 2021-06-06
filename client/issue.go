package client

import (
	"encoding/json"
	"errors"
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

func (g *Gitbucket) SetStatus(hash string, status *Status) error {
	body, err := json.Marshal(status)
	if err != nil {
		log.Println("fail to marshal payload")
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		g.buildStatusURL(hash),
		strings.NewReader(string(body)),
	)
	if err != nil {
		log.Println("fail to create request")
		return err
	}

	req.Header.Set("Authorization", g.buildToken())
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("fail to send request")
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("request do not succee %#v", res))
	}

	return nil
}

func (g *Gitbucket) buildStatusURL(commit string) string {
	return fmt.Sprintf("http://%s/api/v3/repos/%s/%s/statuses/%s", g.host, g.owner, g.repositoryName, commit)
}
