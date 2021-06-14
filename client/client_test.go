package client

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Run("request has authentication header", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", `=~http://githost\:8080/api/v3/`,
			func(req *http.Request) (*http.Response, error) {
				if _, ok := req.Header["Authorization"]; ok {
					res := httpmock.NewStringResponse(200, "")
					return res, nil
				}
				res := httpmock.NewStringResponse(500, "")
				return res, nil
			},
		)

		cli := New("githost:8080", "root", "test")
		cli.SetToken("token")
		res, _ := cli.request("GET", "http://githost:8080/api/v3/", strings.NewReader(""))
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("requestWithoutAuth does not have authentication header", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", `=~http://githost\:8080/api/v3/`,
			func(req *http.Request) (*http.Response, error) {
				if _, ok := req.Header["Authorization"]; ok {
					res := httpmock.NewStringResponse(500, "")
					return res, nil
				}
				res := httpmock.NewStringResponse(200, "")
				return res, nil
			},
		)

		cli := New("githost:8080", "root", "test")
		res, _ := cli.requestWithoutAuth("GET", "http://githost:8080/api/v3/", strings.NewReader(""))
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("EnableHttps set https request", func(t *testing.T) {

		cli := New("githost:8080", "root", "test")
		cli.EnableHttps()
		assert.Equal(t, "https://githost:8080/api/v3", cli.baseUrl())
	})
}
