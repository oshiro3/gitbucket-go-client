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
	log.Printf("%#v\n", comment)
	body, err := json.Marshal(comment)
	log.Printf("%#v\n", string(body))
	if err != nil {
		log.Printf("fail to create request")
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		g.buildCommentURL(issueNumber),
		strings.NewReader(string(body)),
	)
	if err != nil {
		log.Printf("fail to create request")
		return nil, err
	}

	req.Header.Set("Authorization", g.buildToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *Gitbucket) buildCommentURL(number int) string {
	return fmt.Sprintf("http://%s/api/v3/repos/%s/%s/issues/%d/comments", g.host, g.owner, g.repositoryName, number)
}
