package main

import (
	notifier "github.com/AlexsJones/squawker"
	"github.com/AlexsJones/squawker/services/slack"
	"github.com/AlexsJones/squawker/services/stackdriver"
)

func main() {

	notifierManager := notifier.NewManager()

	notifierManager.AddNotifier(&slack.SlackNotifier{ClientToken: "YOURCLIENTOKEN", Channels: []string{"ACHANNEL"}})

	notifierManager.AddNotifier(&stackdriver.StackdriverNotifier{ProjectName: "PROJECT_NAME", LogID: "Logger0"})

	notifierManager.Send("THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")
}
