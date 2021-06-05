package client

import (
	"fmt"
)

type gitbucket struct {
	host       string
	owner      string
	repository string
	token      string
}

func New(host string, owner, repository string) *gitbucket {
	return &gitbucket{
		host:       host,
		owner:      owner,
		repository: repository,
	}
}

func (g *gitbucket) SetToken(token string) {
	g.token = token
}

func (g *gitbucket) buildToken() string {
	return fmt.Sprintf("token %s", g.token)
}
