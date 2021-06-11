package main

import (
	"log"

	"github.com/oshiro3/gitbucket-go-client/client"
)

func main() {
	cli := client.NewClient("192.168.10.2:8080/gitbucket")
	param := &client.UserRepoParam{UserType: "member", PerPage: 2}
	repos, err := cli.GetUserRepositories("Hojo", param)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v\n", repos)
}
