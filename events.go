package rkit

import "sync"

// EventType is the type of event.
type EventType string

// Event is the interface every concrete event implements, and fired from any
// EventSource.∑
type Event interface {
	// String returns the string representation of the event.
	String() string
	// Type ruturns the type of event
	Type() EventType
}

// BaseEvent can be used as base of any event.
type BaseEvent struct {
	name string
}

// String returns string representation of the event.
func (be BaseEvent) String() string {
	return be.name
}

// Type returns the type of events.
func (be BaseEvent) Type() EventType {
	return EventType(be.name)
}

// EventSourceWatcher is an interface any EventSource can implement to be
// notified when there is no one watching or when someone starts watching for
// the first time.
type EventSourceWatcher interface {
	// Start is fired when the first subscriber for this event subscribes.
	// EventSource can decide to start detecting events only after someone is
	// actually listening to it
	Start()
	// Start is fired when the last subscriber has unsubscribed.
	Stop()
}

// EventSource struct, the base struct, (TODO: make this an interface?), that
// any eventsource represents.
type EventSource struct {
	// Channels stores a unique channel for each subscriber. The channel is
	// buffered.
	Channels map[chan Event]struct{}
	// Lock used to make Channels safe. Can also be used by eventsource or
	// EventSourceWatcher implementations for other reasons.
	Lock sync.RWMutex
	// Watcher can be nil if event source is not interested in Start()/Stop()
	// events of EventSourceWatcher
	Watcher EventSourceWatcher
}

// MakeEventSource can be used to create a new event source.
func MakeEventSource() *EventSource {
	return &EventSource{
		Channels: make(map[chan Event]struct{}),
	}
}

// Sub can be used for someone who is interested in an Event to listen for them.
// Returns a Event channel, which is buffered.
//
// When the first guy subscribes, EventSourceWatcher.Start() is called if
// EventSource.Watcher is not nil.
func (s *EventSource) Sub() chan Event {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	ch := make(chan Event, 10)
	s.Channels[ch] = struct{}{}

	if s.Watcher != nil && len(s.Channels) == 1 {
		s.Watcher.Start()
	}
	return ch
}

// Unsub can be used to unsubscribe from an EventSource, eg if you are no longer
// interested in keyboard event. Must be called with the channel returned by
// Sub() call.
//
// When the last guy unsubscribes, EventSourceWatcher.Stop() is called if
// EventSource.Watcher is not nil.
func (s *EventSource) Unsub(ch chan Event) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	delete(s.Channels, ch)
	if s.Watcher != nil && len(s.Channels) == 0 {
		s.Watcher.Stop()
	}
}

// Pub is used to publish an event to all the subscribers of an eventsource.
func (s *EventSource) Pub(ev Event) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	/*
		TODO: how to decide if this is a safe pattern? can one lose some events
		if they are busy processing stuff? i suppose if they really care, they
		can spawn a goroutine dedicated to capturing all fired events, and que
		them elsewhere and process them later.
	*/
	for ch := range s.Channels {
		select {
		case ch <- ev:
		default:
		}
	}
}
