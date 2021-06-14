package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestRepositories(t *testing.T) {
	t.Run("getPublicRepositories return repositories", func(t *testing.T) {
		resp := `[{"name":"test","full_name":"root/test","description":"","watchers":0,"forks":0,"private":false,"default_branch":"master","owner":{"login":"root","email":"root@localhost","type":"User","site_admin":true,"created_at":"2021-06-05T09:13:25Z","id":0,"url":"http://localhost:8081/api/v3/users/root","html_url":"http://localhost:8081/root","avatar_url":"http://localhost:8081/root/_avatar"},"has_issues":true,"id":0,"forks_count":0,"watchers_count":0,"url":"http://localhost:8081/api/v3/repos/root/test","clone_url":"http://localhost:8081/git/root/test.git","html_url":"http://localhost:8081/root/test"}]`
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", `=~http://githost\:8080/api/v3/repositories`,
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, resp)
				return res, nil
			},
		)

		res, err := NewClient("githost:8080").GetPublicRepositories()

		assert.NoError(t, err)
		assert.Equal(t, "test", res[0].Name)
	})

	t.Run("getUserRepositories return repositories", func(t *testing.T) {
		resp := `[{"name":"test","full_name":"root/test","description":"","watchers":0,"forks":0,"private":false,"default_branch":"master","owner":{"login":"root","email":"root@localhost","type":"User","site_admin":true,"created_at":"2021-06-05T09:13:25Z","id":0,"url":"http://localhost:8081/api/v3/users/root","html_url":"http://localhost:8081/root","avatar_url":"http://localhost:8081/root/_avatar"},"has_issues":true,"id":0,"forks_count":0,"watchers_count":0,"url":"http://localhost:8081/api/v3/repos/root/test","clone_url":"http://localhost:8081/git/root/test.git","html_url":"http://localhost:8081/root/test"}]`
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", `=~http://githost\:8080/api/v3/users/*`,
			func(req *http.Request) (*http.Response, error) {
				res := httpmock.NewStringResponse(200, resp)
				return res, nil
			},
		)

		param := &UserRepoParam{UserType: "member"}
		res, err := NewClient("githost:8080").GetUserRepositories("user", param)

		assert.NoError(t, err)
		assert.Equal(t, "test", res[0].Name)
	})
}
