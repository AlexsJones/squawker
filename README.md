# squawker

A plug n play notifications system for golang

```go
	notifierManager := notifier.NewManager()

	notifierManager.AddNotifier(&slack.SlackNotifier{ClientToken: "YOURCLIENTOKEN", Channels: []string{"ACHANNEL"}})

	notifierManager.AddNotifier(&stackdriver.StackdriverNotifier{ProjectName: "PROJECT_NAME", LogID: "Logger0"})

	notifierManager.Send("THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")
```
