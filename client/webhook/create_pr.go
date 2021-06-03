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
	Owner         user   `json:"owner"`
	Id            int    `json:"id"`
	ForksCount    int    `json:"forks_count"`
	WatchersCount int    `json:"watchers_count"`
	Url           string `json:"url"`
	HttpUrl       string `json:"http_url"`
	CloneUrl      string `json:"clone_url"`
	HtmlUrl       string `json:"html_url"`
	SshUrl        string `json:"ssh_url"`
}

type user struct {
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

type base struct {
	Sha  string     `json:"sha"`
	Ref  string     `json:"ref"`
	Repo repository `json:"repo"`
}

type head struct {
	Sha  string     `json:"sha"`
	Ref  string     `json:"ref"`
	Repo repository `json:"repo"`
	//TODO
	// Label
	User user `json:"user"`
}

type pullReauest struct {
	Number    int    `json:"repository"`
	State     string `json:"state"`
	UpdatedAt string `json:"updatead_at"`
	CreatedAt string `json:"created_at"`
	Head      head   `json:"head"`
	Base      base   `json:"base"`
	Merged    bool   `json:"merged"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	User      user   `json:"user"`
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

type PRWebhook struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	Repository  repository  `json:"repository"`
	PullReauest pullReauest `json:"pull_request"`
	Sender      user        `json:"sender"`
}

var (
	OPEN  string = "opened"
	CLOSE string = "closed"
)

func ParsePRWebhook(body io.Reader) *PRWebhook {
	var h *PRWebhook
	err := json.NewDecoder(body).Decode(&h)
	if err != nil {
		log.Fatal(err)
	}
	return h
}

func (h *CreatePRWebhook) IsOpened() bool {
	return h.Action == OPEN
}

func (h *CreatePRWebhook) IsClosed() bool {
	return h.Action == CLOSE
}

func (h *CreatePRWebhook) IsMerged() bool {
	return h.IsClosed() && h.PullReauest.Merged
}
