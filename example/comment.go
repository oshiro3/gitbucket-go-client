package main

import (
	"log"

	"github.com/oshiro3/gitbucket-go-client/client"
)

func main() {
	cli := client.New("127.0.0.1:8081", "root", "test")
	cli.SetToken(os.Getenv("token"))
	res, err := cli.Comment(1, "test")
	if err != nil {
		log.Printf("%#v\n", err)
	}
	log.Printf("res: %#v\n", res)

}
