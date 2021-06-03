package main

import (
	"bytes"
	"io/ioutil"
	"log"

	client "github.com/oshiro3/gitbucket-go-client/client/webhook"
)

func main() {
	data, err := ioutil.ReadFile("res")
	if err != nil {
		panic(err)
	}
	log.Println(string(data[8:]))
	client.ParseCreatePRWebhook(bytes.NewReader(data[8:]))
}
