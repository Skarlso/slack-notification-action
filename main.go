package main

import (
	"log"

	"github.com/Skarlso/slack-notification-action/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
