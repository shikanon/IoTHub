package hub

type Notifier interface {
	SentEvent() error
}
