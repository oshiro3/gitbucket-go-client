package client

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type gitbucket struct {
	host       string
	port       int
	owner      string
	repository string
	token      string
}

func New(host string, port int, owner, repository string) *gitbucket {
	return &gitbucket{
		host:       host,
		port:       port,
		owner:      owner,
		repository: repository,
	}
}

type comment struct {
	body string `json: "body"`
}

func (g *gitbucket) Comment(issueNumber int, body string) *http.Response {
	req, err := http.NewRequest(
		http.MethodPost,
		g.buildCommentURL(issueNumber),
		strings.NewReader(fmt.Sprintf("{\"body\": \"%s\"}", body)),
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", g.buildToken())
	log.Printf("%v\n", req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", res)
	return res
}

func (g *gitbucket) SetToken(token string) {
	g.token = token
}

func (g *gitbucket) buildToken() string {
	return fmt.Sprintf("token %s", g.token)
}

func (g *gitbucket) buildCommentURL(number int) string {
	return fmt.Sprintf("http://%s:%d/api/v3/repos/%s/%s/issues/%d/comments", g.host, g.port, g.owner, g.repository, number)
}
