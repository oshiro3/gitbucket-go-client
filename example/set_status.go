package main

import (
	"log"

	"github.com/oshiro3/gitbucket-go-client/client"
)

func main() {
	cli := client.New("127.0.0.1:8081", "root", "test")
	cli.SetToken(os.Getenv("token")
	payload := &client.Status{State: "pending", TargetUrl: "http://hogehoge.com", Description: "test message", Context: "ci"}
	err := cli.SetStatus("$commit_hash", payload)
	if err != nil {
		log.Printf("%#v\n", err)
	}
}
