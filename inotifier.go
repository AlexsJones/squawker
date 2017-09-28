package notifier

//INotifier ...
type INotifier interface {
	Create() error
	GetName() string
	Notify(string, ...string) error
	Destroy()
}
