package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/go-playground/webhooks.v5/github"
)

// Config WEBHOOK_PORT, WEBHOOK_SECRET
type Config struct {
	Port                 string `default:"3000"`
	Secret               string `required:"true"`
	DeploymentBranch     string `required:"true" split_words:"true"`
	DeploymentScriptPath string `required:"true" split_words:"true"`
}

const (
	path = "/webhooks"
)

func deploy(c Config) {
	log.Println("deployment webhook start:", c.DeploymentBranch)
	out, err := exec.Command(c.DeploymentScriptPath).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	if out != nil {
		log.Println(string(out))
	}
	log.Println("deployment webhook finished")
}

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
				go deploy(c)
			}
		}
	})

	addr := ":" + c.Port
	log.Println("listen on", addr)
	http.ListenAndServe(addr, nil)
}
