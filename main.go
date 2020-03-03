package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret(os.Getenv("WEBHOOK_SECRET")))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				fmt.Printf("%+v", err)
			}
		}
		switch payload.(type) {
		case github.PushPayload:
			push := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", push)
			fmt.Print(push.Ref)
			fmt.Print(push.BaseRef)
		}
	})
	http.ListenAndServe(":3000", nil)
}
