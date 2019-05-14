package event

import (
	"time"
)

// Message represents a message event
type Message interface {
	Key() string
}

// MeowCreatedMessage is the event created when the Meow is created
type MeowCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

// Key returns the key of event created
func (m *MeowCreatedMessage) Key() string {
	return "meow.created"
}
