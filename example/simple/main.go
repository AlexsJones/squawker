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

	var notifiers = []notifier.INotifier{
		&stackdriver.StackdriverNotifier{ProjectName: "PROJECT", LogID: "Logger0"},
		&slack.SlackNotifier{ClientToken: "xoxp-2310897947-180906251303-247965629927-81a3b313ae73ad0f7c4c16c8a835b1c9", Channels: []string{"CHANNEL"}},
		&victorops.VictorOpsNotifier{APIKey: "RestfulIntegrationKeyxxxx", RoutingKey: "production",
			EntityID: "XXX", MessageStatus: vo.Warning},
	}

	for _, noti := range notifiers {
		notifierManager.AddNotifier(noti)
	}

	notifierManager.SendRoutes("Warning", "THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")
}
