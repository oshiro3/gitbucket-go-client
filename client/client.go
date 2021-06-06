package client

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Repository struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	Watchers      int    `json:"watchers"`
	Forks         int    `json:"forks"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
	Owner         User   `json:"owner"`
	Id            int    `json:"id"`
	ForksCount    int    `json:"forks_count"`
	WatchersCount int    `json:"watchers_count"`
	Url           string `json:"url"`
	HttpUrl       string `json:"http_url"`
	CloneUrl      string `json:"clone_url"`
	HtmlUrl       string `json:"html_url"`
	SshUrl        string `json:"ssh_url"`
}

type User struct {
	Login     string `json:"login"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Url       string `json:"url"`
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}

type Gitbucket struct {
	owner          string
	repositoryName string
	host           string
	token          string
	https          bool
}

func New(host string, owner, repositoryName string) *Gitbucket {
	return &Gitbucket{
		owner:          owner,
		repositoryName: repositoryName,
		host:           host,
		https:          false,
	}
}

func NewClient(host string) *Gitbucket {
	return &Gitbucket{
		host:  host,
		https: false,
	}
}

func (g *Gitbucket) SetToken(token string) {
	g.token = token
}

func (g *Gitbucket) EnableHttps() {
	g.https = true
}

func (g *Gitbucket) buildToken() string {
	return fmt.Sprintf("token %s", g.token)
}

func (g *Gitbucket) baseUrl() string {
	var protocol string
	if g.https {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s/api/v3", protocol, g.host)
}

func (g *Gitbucket) request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		log.Println("fail to create request")
		return nil, err
	}

	req.Header.Set("Authorization", g.buildToken())
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("fail to send request")
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("request do not succee %#v", res))
	}
	return res, nil
}
