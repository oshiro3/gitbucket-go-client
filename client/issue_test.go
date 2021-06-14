package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSetStatus(t *testing.T) {
	t.Run("setStatus return nil if request succeed", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", `=~http://githost\:8080/api/v3/repos/root/test/statuses/abc`,
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, "")
				return res, nil
			},
		)

		cli := New("githost:8080", "root", "test")
		payload := &Status{State: "pending", TargetUrl: "http://hogehoge.com", Description: "test message", Context: "ci"}
		res, err := cli.SetStatus("abc", payload)
		assert.Equal(t, 200, res.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("buildStatusURL return gitbucket status url", func(t *testing.T) {
		cli := New("githost:8080", "root", "test")
		expect := "http://githost:8080/api/v3/repos/root/test/statuses/abc"
		assert.Equal(t, expect, cli.buildStatusURL("abc"))

	})

	t.Run("return gitbucket status url", func(t *testing.T) {
		cli := New("githost:8080", "root", "test")
		expect := "http://githost:8080/api/v3/repos/root/test/statuses/abc"
		assert.Equal(t, expect, cli.buildStatusURL("abc"))
	})
}
