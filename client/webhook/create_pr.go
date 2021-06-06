package webhook

import (
	"encoding/json"
	"io"
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

func parsePRWebhook(body io.Reader) (*PRWebhook, error) {
	var h *PRWebhook
	err := json.NewDecoder(body).Decode(&h)
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
