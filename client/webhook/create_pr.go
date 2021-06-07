package webhook

import (
	"encoding/json"
	"net/http"
	"strings"
)

type PRWebhook struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pull_request"`
	Sender      User        `json:"sender"`
}

var (
	OPEN  string = "opened"
	CLOSE string = "closed"
	SYNC  string = "synchronize"
)

func ParsePRWebhook(r *http.Request) (*PRWebhook, error) {
	var h *PRWebhook
	payload := r.FormValue("payload")
	err := json.NewDecoder(strings.NewReader(payload)).Decode(&h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h *PRWebhook) IsOpened() bool {
	return h.Action == OPEN
}

func (h *PRWebhook) IsSynchronize() bool {
	return h.Action == SYNC
}

func (h *PRWebhook) IsClosed() bool {
	return h.Action == CLOSE
}

func (h *PRWebhook) IsMerged() bool {
	return h.IsClosed() && h.PullRequest.Merged
}
