package notifier

import "errors"

//Manager ...
type Manager struct {
	Notifiers map[string]inotifier
}

//NewManager creates a new notifier Manager
func NewManager() *Manager {

	return &Manager{Notifiers: make(map[string]inotifier)}
}

//AddNotifier adds a notifier to the Manager
func (s *Manager) AddNotifier(i inotifier) error {
	s.Notifiers[i.GetName()] = i
	//Calls create for any additional setup
	if err := i.Create(); err != nil {
		return err
	}
	return nil
}

//SendFanOut a message to all notifiers
func (s *Manager) SendFanOut(d ...string) {

	for _, value := range s.Notifiers {
		value.Notify(d...)
	}
}

//SendFanIn a message to a single notifier
func (s *Manager) SendFanIn(notifierName string, d ...string) error {

	if value, err := s.Notifiers[notifierName]; err != false {
		return value.Notify(d...)
	}
	return errors.New("Unable to find notifier")
}
