package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type comment struct {
	Body string `json:"body"`
}

func (g *Gitbucket) Comment(issueNumber int, message string) (*http.Response, error) {
	comment := &comment{Body: message}
	body, err := json.Marshal(comment)
	if err != nil {
		log.Printf("fail to create request")
		return nil, err
	}

	res, err := g.request(
		"POST",
		g.buildCommentURL(issueNumber),
		strings.NewReader(string(body)),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *Gitbucket) buildCommentURL(number int) string {
	return fmt.Sprintf("%s/repos/%s/%s/issues/%d/comments", g.baseUrl(), g.owner, g.repositoryName, number)
}
