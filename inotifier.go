package notifier

type inotifier interface {
	Create() error
	GetName() string
	Notify(...string) error
	Destroy()
}
