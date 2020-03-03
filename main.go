package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/go-playground/webhooks.v5/github"
)

// Config WEBHOOK_PORT, WEBHOOK_SECRET
type Config struct {
	Port             string `default:"3000"`
	Secret           string `required:"true"`
	DeploymentBranch string `required:"true" split_words:"true"`
}

const (
	path = "/webhooks"
)

func main() {
	var c Config
	err := envconfig.Process("webhook", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	hook, _ := github.New(github.Options.Secret(c.Secret))

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
			if push.Ref == "refs/heads/"+c.DeploymentBranch {
				fmt.Printf("%+v\n", push)
			}
			fmt.Println(push.BaseRef)
		}
	})

	addr := ":" + c.Port
	log.Println("listen on", addr)
	http.ListenAndServe(addr, nil)
}
