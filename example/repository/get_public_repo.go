package main

import (
	"log"

	"github.com/oshiro3/gitbucket-go-client/client"
)

func main2() {
	cli := client.NewClient("localhost:8081")
	repos, err := cli.GetPublicRepositories()
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v\n", repos)
}
