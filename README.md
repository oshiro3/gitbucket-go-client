# gitbucket-go-client

`gitbucket-go-client` is a library for manipulating gitbucket by calling the gitbucket API from a go program.

## getting started

```
go get -u github.com/oshiro3/gitbucket-go-client
```

## usage

```
import (
	"os"

	"github.com/oshiro3/gitbucket-go-client/client"
)

func main() {
	cli := client.New(gitbucketHost, repositoryOwner, repositoryName)
	cli.SetToken(os.Getenv("token"))
	payload := &client.Status{State: "pending", TargetUrl: "http://example.com", Description: "example message", Context: "ci"}
	err := cli.SetStatus("$commit_hash", payload)
}
```

## caution

I haven't versioned yet because it's in the early stages of development.

