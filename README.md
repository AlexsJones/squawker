# squawker

[![Maintainability](https://api.codeclimate.com/v1/badges/b1112f79f38545b8a831/maintainability)](https://codeclimate.com/github/AlexsJones/squawker/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlexsJones/squawker)](https://goreportcard.com/report/github.com/AlexsJones/squawker)
A plug n play notifications system for golang


### Example

```go
func main() {

	notifierManager := notifier.NewManager()

	var notifiers = []notifier.INotifier{
		&stackdriver.StackdriverNotifier{ProjectName: "PROJECT", LogID: "Logger0"},
		&slack.SlackNotifier{ClientToken: "xoxp-TOKENNAME", Channels: []string{"CHANNEL"}},
		&victorops.VictorOpsNotifier{APIKey: "RestfulIntegrationKeyxxxx", RoutingKey: "production",
			EntityID: "XXX", MessageStatus: vo.Warning},
	}

	for _, noti := range notifiers {
		if err := notifierManager.AddNotifier(noti); err != nil {
			log.Fatal(err)
		}
	}

	notifierManager.SendFanOut("Warning", "THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")
}

```

### Example using routing keys

```go
package main

import (
	notifier "github.com/AlexsJones/squawker"
	"github.com/AlexsJones/squawker/services/slack"
	"github.com/AlexsJones/squawker/services/stackdriver"
	"github.com/AlexsJones/squawker/services/victorops"
	vo "github.com/chrissnell/victorops-go"
)

func main() {

	notifierManager := notifier.NewManager()

	var notifiers = map[string]notifier.INotifier{
		"Log":     &stackdriver.StackdriverNotifier{ProjectName: "PROJECT", LogID: "Logger0"},
		"Warning": &slack.SlackNotifier{ClientToken: "xoxp-TOKENNAME", Channels: []string{"CHANNEL"}},
		"Critical": &victorops.VictorOpsNotifier{APIKey: "RestfulIntegrationKeyxxxx", RoutingKey: "production",
			EntityID: "XXX", MessageStatus: vo.Warning},
	}

	notifierManager.AddNotifiersWithRoutingKeys(notifiers)

	notifierManager.SendFanOut("Warning", "This is an event some of you need to know about!")
}
```
