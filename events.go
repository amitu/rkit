package rkit

import "sync"

type EventType string

type Event interface {
	String() string
	Type() EventType
}

type BaseEvent struct {
	name string
}

func (be BaseEvent) String() string {
	return be.name
}

func (be BaseEvent) Type() EventType {
	return EventType(be.name)
}

type EventSource struct {
	channels map[chan Event]struct{}
	lock     sync.RWMutex
}

func MakeEventSource() *EventSource {
	return &EventSource{
		channels: make(map[chan Event]struct{}),
	}
}

func (s *EventSource) Sub() chan Event {
	s.lock.Lock()
	defer s.lock.Unlock()

	ch := make(chan Event)
	s.channels[ch] = struct{}{}
	return ch
}

func (s *EventSource) Unsub(ch chan Event) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.channels, ch)
}

func (s *EventSource) Pub(ev Event) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	/*
		TODO: how to decide if this is a safe pattern? can one lose some events
		if they are busy processing stuff? i suppose if they really care, they
		can spawn a goroutine dedicated to capturing all fired events, and que
		them elsewhere and process them later.
	*/
	for ch, _ := range s.channels {
		select {
		case ch <- ev:
		default:
		}
	}
}
