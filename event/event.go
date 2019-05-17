package event

import (
	"github.com/leogsouza/meower/schema"
)

// EventStore is an interface for store events
type EventStore interface {
	Close()
	PublishMeowCreated(meow schema.Meow) error
	SubscribeMeowCreated() (<-chan MeowCreatedMessage, error)
	OnMeowCreated(f func(MeowCreatedMessage)) error
}

var impl EventStore

// SetEventStore sets the EventStore
func SetEventStore(es EventStore) {
	impl = es
}

// Close closes the EventStore
func Close() {
	impl.Close()
}

// PublishMeowCreated send an event to all subscribers
func PublishMeowCreated(meow schema.Meow) error {
	return impl.PublishMeowCreated(meow)
}

// SubscribeMeowCreated adds a new subscriber for the event
func SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	return impl.SubscribeMeowCreated()
}

// OnMeowCreated executes a function when MeowCreated
func OnMeowCreated(f func(MeowCreatedMessage)) error {
	return impl.OnMeowCreated(f)
}
