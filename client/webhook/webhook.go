package webhook

import (
	"errors"
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

type Base struct {
	Sha  string     `json:"sha"`
	Ref  string     `json:"ref"`
	Repo Repository `json:"repo"`
}

type Head struct {
	Sha  string     `json:"sha"`
	Ref  string     `json:"ref"`
	Repo Repository `json:"repo"`
	//TODO
	// Label
	User User `json:"user"`
}

type PullRequest struct {
	Number    int    `json:"number"`
	State     string `json:"state"`
	UpdatedAt string `json:"updatead_at"`
	CreatedAt string `json:"created_at"`
	Head      Head   `json:"head"`
	Base      Base   `json:"base"`
	Merged    bool   `json:"merged"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	User      User   `json:"user"`
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

func judgeEvent(r *http.Request) string {
	return r.Header.Get("X-Github-Event")
}

func ParseWebhook(r *http.Request) (interface{}, error) {
	switch judgeEvent(r) {
	case "pull_request":
		return parsePRWebhook(r)
	}
	return nil, errors.New("event type is undefined")
}
