package victorops

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	victorops "github.com/chrissnell/victorops-go"
)

//VictorOpsNotivier ...
type VictorOpsNotifier struct {
	APIKey        string
	RoutingKey    string
	EntityID      string
	MessageStatus victorops.MessageType
	api           *victorops.API
}

//Create a notifier
func (v *VictorOpsNotifier) Create() error {
	if v.APIKey == "" {
		return errors.New("Requires APIKey")
	}
	v.api = victorops.NewClient(v.APIKey)
	return nil
}

//GetName of the notifier
func (v *VictorOpsNotifier) GetName() string {

	return "VictorOps"
}

//Notify ...
func (v *VictorOpsNotifier) Notify(routingKey string, msg ...string) (err error) {
	// Fix until https://github.com/chrissnell/victorops-go/pull/1 is merged in
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	log.Println("Starting VictorOps notify")
	if v.EntityID == "" || v.RoutingKey == "" || v.APIKey == "" {
		return errors.New("Requires EntityID, RoutingKey and APIKey")
	}
	log.Println("Preparing VictorOps payload")
	e := &victorops.Event{
		RoutingKey:        v.RoutingKey,
		MessageType:       v.MessageStatus,
		EntityID:          v.EntityID,
		StateMessage:      strings.Join(msg, " "),
		Timestamp:         time.Now(),
		EntityDisplayName: "Squawker",
	}
	resp, err := v.api.SendAlert(e)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println(resp)
	log.Println("Notified VictorOps")
	return nil
}

//Destroy ...
func (v *VictorOpsNotifier) Destroy() {
}
