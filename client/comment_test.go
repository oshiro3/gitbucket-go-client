package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestComment(t *testing.T) {
	t.Run("comment return nil and http response if request succeed", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", `=~http://githost\:8080/api/v3/repos/root/test/issues/\d/comments`,
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, "")
				return res, nil
			},
		)

		cli := New("githost:8080", "root", "test")
		res, err := cli.Comment(1, "body")
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("buildCommentURL return gitbucket comment url", func(t *testing.T) {
		cli := New("githost:8080", "root", "test")
		expect := "http://githost:8080/api/v3/repos/root/test/issues/1/comments"
		assert.Equal(t, expect, cli.buildCommentURL(1))

	})
}
