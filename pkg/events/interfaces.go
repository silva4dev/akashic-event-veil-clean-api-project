package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHanlderInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHanlderInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHanlderInterface) error
	Has(eventName string, handler EventHanlderInterface) bool
	Clear() error
}
