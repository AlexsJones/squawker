package slack

import (
	"errors"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

//SlackNotifier ...
type SlackNotifier struct {
	Client      *slack.Client
	ClientToken string
	Channels    []string
}

//Create setup functionality of the notifier
func (s *SlackNotifier) Create() error {

	s.Client = slack.New(s.ClientToken)

	if s.Client == nil {
		return errors.New("Unable to create slack client")
	}

	_, err := s.Client.AuthTest()
	return err
}

//GetName of the notifier
func (s *SlackNotifier) GetName() string {
	return "Slack"
}

//Notify the notifier target
func (s *SlackNotifier) Notify(msg ...string) error {
	attachments := slack.Attachment{
		Pretext: "Important Notification",
		Text:    strings.Join(msg, "\n"),
	}
	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{attachments},
	}
	log.Println("The number of channel being messaged are:", len(s.Channels))
	for _, channel := range s.Channels {
		_, _, err := s.Client.PostMessage(channel, "Go Cron Notification", params)
		if err != nil {
			return err
		}
		log.Println("Notified slack channel:", channel)
	}
	return nil
}

//Destroy any teardown code here
func (s *SlackNotifier) Destroy() {

}
