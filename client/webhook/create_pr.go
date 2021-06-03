package client

import (
	"encoding/json"
	"io"
	"log"
)

type repository struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	Watchers      int    `json:"watchers"`
	Forks         int    `json:"forks"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
	//TODO
	// Owner         user   `json:"full_name"`
	Id            int    `json:"id"`
	ForksCount    int    `json:"forks_count"`
	WatchersCount int    `json:"watchers_count"`
	Url           string `json:"url"`
	HttpUrl       string `json:"http_url"`
	CloneUrl      string `json:"clone_url"`
	HtmlUrl       string `json:"html_url"`
	SshUrl        string `json:"ssh_url"`
}

type base struct {
	Sha  string     `json:"sha"`
	Ref  string     `json:"ref"`
	Repo repository `json:"repo"`
}

type pullReauest struct {
	Number    int    `json:"repository"`
	State     string `json:"state"`
	UpdatedAt string `json:"updatead_at"`
	CreatedAt string `json:"created_at"`
	//TODO
	// Head
	Base   base   `json:"base"`
	Merged bool   `json:"merged"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	//TODO
	// User
	//TODO
	// Labels            []label `json:"labels"`
	Id                int    `json:"id"`
	CommitsUrl        string `json:"commits_url"`
	ReviewCommentsUrl string `json:"reveiw_comments_url"`
	ReviewCommentUrl  string `json:"review_comment_url"`
	CommentsUrl       string `json:"comments_url"`
	StatusesUrl       string `json:"statuses_url"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
}

type CreatePRWebhook struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	Repository  repository  `json:"repository"`
	PullReauest pullReauest `json:"pull_request"`
	//TODO
	// Sender sender `json:"sender"`
}

func ParseCreatePRWebhook(body io.Reader) CreatePRWebhook {
	var h CreatePRWebhook
	err := json.NewDecoder(body).Decode(&h)
	if err != nil {
		log.Fatal(err)
	}
	return h
}
