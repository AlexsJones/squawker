package notifier

import (
	"errors"
	"fmt"
	"log"
)

//Manager ...
type Manager struct {
	notifiers       map[string]INotifier
	routingkeyRules map[string][]string
	errorHandler    func(notifier INotifier, err error)
}

//NewManager creates a new notifier Manager
func NewManager(errorHandler func(notifier INotifier, err error)) *Manager {

	return &Manager{notifiers: make(map[string]INotifier), routingkeyRules: make(map[string][]string),
		errorHandler: errorHandler}
}

//AddNotifier adds a notifier to the Manager
func (s *Manager) AddNotifier(i INotifier) {
	s.notifiers[i.GetName()] = i
	//Calls create for any additional setup
	if err := i.Create(); err != nil {
		s.errorHandler(i, err)
	}
}

//AddNotifierWithRoutingKey for selecting notifier based on event type in a fan out operation
func (s *Manager) AddNotifierWithRoutingKey(routingKey string, n INotifier) {

	err := n.Create()
	s.notifiers[n.GetName()] = n
	s.routingkeyRules[routingKey] = append(s.routingkeyRules[routingKey], n.GetName())
	log.Println(fmt.Sprintf("Adding Routing key: %s for %s\n", routingKey, n.GetName()))
	if s.errorHandler != nil && err != nil {
		s.errorHandler(n, err)
	}
}

//SendRoutes a message to all notifiers
func (s *Manager) SendRoutes(routingkey string, d ...string) {

	if v, ok := s.routingkeyRules[routingkey]; ok != false {

		for _, arrItem := range v {
			if err := s.SendFanIn(arrItem, routingkey, d...); err != nil {
				if s.errorHandler != nil {
					s.errorHandler(s.notifiers[arrItem], err)
				}
			}
		}
	}
}

//SendFanIn a message to a single notifier
func (s *Manager) SendFanIn(notifierName string, routingKey string, d ...string) error {

	if value, err := s.notifiers[notifierName]; err != false {
		return value.Notify(routingKey, d...)
	}
	return errors.New("Unable to find notifier")
}

// Closenotifiers will go through all the added notifiers and call their
// destroy method and remove them from the NotifyManager
func (s *Manager) Closenotifiers() {
	for key, value := range s.notifiers {
		log.Println(fmt.Sprintf("Removing %v from the notifier map", key))
		value.Destroy()
		delete(s.notifiers, key)
	}
}
