package main

import (
	"log"

	notifier "github.com/AlexsJones/squawker"
	"github.com/AlexsJones/squawker/services/slack"
	"github.com/AlexsJones/squawker/services/stackdriver"
	"github.com/AlexsJones/squawker/services/victorops"
	vo "github.com/chrissnell/victorops-go"
)

func main() {

	notifierManager := notifier.NewManager(func(notifier notifier.INotifier, err error) {
		log.Fatal(err.Error())
	})

	notifierManager.AddNotifierWithRoutingKey("Log", &stackdriver.StackdriverNotifier{ProjectName: "PROJECT", LogID: "Logger0"})
	notifierManager.AddNotifierWithRoutingKey("Warning", &slack.SlackNotifier{ClientToken: "xoxp-2XX", Channels: []string{"XXXX"}})
	notifierManager.AddNotifierWithRoutingKey("Warning", &victorops.VictorOpsNotifier{APIKey: "a49f43XXXX", RoutingKey: "production",
		EntityID: "XXXX", MessageStatus: vo.Critical})
	notifierManager.SendRoutes("Warning", "THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")

}
