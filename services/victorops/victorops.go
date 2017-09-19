package victorops

import (
	"errors"
	victorops "github.com/chrissnell/victorops-go"
	"log"
	"strings"
	"time"
)

//VictorOpsNotivier ...
type VictorOpsNotivier struct {
	APIKey        string
	RoutingKey    string
	EntityID      string
	MessageStatus victorops.MessageType
	api           *victorops.API
}

//Create a notifier
func (v *VictorOpsNotivier) Create() error {
	v.api = victorops.NewClient(v.APIKey)
	return nil
}

//GetName of the notifier
func (v *VictorOpsNotivier) GetName() string {

	return "VictorOps"
}

//Notify ...
func (v *VictorOpsNotivier) Notify(msg ...string) error {
	if v.EntityID == "" || v.RoutingKey == "" || v.APIKey == "" {
		return errors.New("Requires EntityID, RoutingKey and APIKey")
	}
	e := &victorops.Event{
		RoutingKey:        v.RoutingKey,
		MessageType:       v.MessageStatus,
		EntityID:          v.EntityID,
		StateMessage:      strings.Join(msg, " "),
		Timestamp:         time.Now(),
		EntityDisplayName: "Squawker",
	}
	log.Println(e)
	resp, err := v.api.SendAlert(e)
	if err != nil {
		return err
	}
	log.Println(resp)

	return nil
}

//Destroy ...
func (v *VictorOpsNotivier) Destroy() {
}
