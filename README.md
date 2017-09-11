# squawker

A plug n play notifications system for golang

```go
	notifierManager := notifier.NewManager()

	if err := notifierManager.AddNotifier(&slack.SlackNotifier{ClientToken: "YOURCLIENTOKEN", Channels: []string{"ACHANNEL"}}); err != nil {
    log.Fatal(err)
  }

	if err := notifierManager.AddNotifier(&stackdriver.StackdriverNotifier{ProjectName: "PROJECT_NAME", LogID: "Logger0"}); err != nil {
    log.Fatal(err)
  }

	notifierManager.SendFanOut("THIS IS EVENT YOU ALL NEED TO KNOW ABOUT!")

  notifierManager.SendFanIn("Slack","An alert has been sent...")
```
