package main

import (
	"log"

	notifier "github.com/AlexsJones/squawker"
	"github.com/AlexsJones/squawker/services/slack"
	"github.com/AlexsJones/squawker/services/stackdriver"
)

func main() {

	notifierManager := notifier.NewManager()

	if err := notifierManager.AddNotifier(&slack.SlackNotifier{ClientToken: "xoxp-TOKENNAME", Channels: []string{"CHANNEL"}}); err != nil {
		log.Fatal(err)

	}

	if err := notifierManager.AddNotifier(&stackdriver.StackdriverNotifier{ProjectName: "PROJECT", LogID: "Logger0"}); err != nil {
		log.Fatal(err)
	}

	notifierManager.SendFanOut("THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")
}
