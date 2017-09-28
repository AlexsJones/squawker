package stackdriver

import (
	"context"
	"errors"
	"log"
	"strings"

	"cloud.google.com/go/logging"
)

//StackdriverNotifier ...
type StackdriverNotifier struct {
	ProjectName string
	LogID       string
	client      *logging.Client
	logger      *logging.Logger
}

//Create a notifier
func (s *StackdriverNotifier) Create() error {
	ctx := context.Background()

	client, err := logging.NewClient(ctx, s.ProjectName)
	if err != nil {
		return err
	}
	s.client = client
	s.logger = client.Logger(s.LogID)

	return nil
}

//GetName for the notifier
func (s *StackdriverNotifier) GetName() string {
	return "Stackdriver"
}

//Notify the notifier
func (s *StackdriverNotifier) Notify(routingKey string, msg ...string) error {
	log.Println("Starting Stackdriver notify")
	if s.ProjectName == "" || s.LogID == "" {
		return errors.New("Requires both ProjectName and LogID parameters")
	}
	log.Println("Creating payload")
	entry := logging.Entry{Payload: strings.Join(msg, " ")}
	s.logger.Log(entry)
	s.logger.Flush()
	log.Println("Notified Stackdriver")
	return nil
}

//Destroy ...
func (s *StackdriverNotifier) Destroy() {
	s.client.Close()
}
