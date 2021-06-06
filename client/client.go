package client

import (
	"fmt"
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
}

func New(host string, owner, repositoryName string) *Gitbucket {
	return &Gitbucket{
		owner:          owner,
		repositoryName: repositoryName,
		host:           host,
	}
}

func NewClient(host string) *Gitbucket {
	return &Gitbucket{
		host: host,
	}
}

func (g *Gitbucket) SetToken(token string) {
	g.token = token
}

func (g *Gitbucket) buildToken() string {
	return fmt.Sprintf("token %s", g.token)
}
