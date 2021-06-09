package main

import (
	"log"
	"net/http"

	webhook "github.com/oshiro3/gitbucket-go-client/client/webhook"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	event := webhook.JudgeEvent(r)
	switch event {
	case "pull_request":
		hook, err := webhook.ParsePRWebhook(r)
		if err != nil {
			log.Printf("Fail to parse webhook: %v\n", err)
			return
		}
		log.Printf("%v\n", webhook)
	}
}

func main() {
	http.HandleFunc("/", sampleHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
