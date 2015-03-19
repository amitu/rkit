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

type EventSourceWatcher interface {
	Start()
	Stop()
}

type EventSource struct {
	Channels map[chan Event]struct{}
	Lock     sync.RWMutex
	Watcher  EventSourceWatcher
}

func MakeEventSource() *EventSource {
	return &EventSource{
		Channels: make(map[chan Event]struct{}),
	}
}

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

func (s *EventSource) Unsub(ch chan Event) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	delete(s.Channels, ch)
	if s.Watcher != nil && len(s.Channels) == 0 {
		s.Watcher.Stop()
	}
}

func (s *EventSource) Pub(ev Event) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	/*
		TODO: how to decide if this is a safe pattern? can one lose some events
		if they are busy processing stuff? i suppose if they really care, they
		can spawn a goroutine dedicated to capturing all fired events, and que
		them elsewhere and process them later.
	*/
	for ch, _ := range s.Channels {
		select {
		case ch <- ev:
		default:
		}
	}
}
