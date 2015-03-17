package rkit

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
}

func MakeEventSource() *EventSource {
	return &EventSource{
		channels: make(map[chan Event]struct{}),
	}
}

func (s *EventSource) Sub() chan Event {
	ch := make(chan Event)
	s.channels[ch] = struct{}{}
	return ch
}

func (s *EventSource) Unsub(ch chan Event) {
	delete(s.channels, ch)
}

func (s *EventSource) Pub(ev Event) {
	for ch, _ := range s.channels {
		select {
		case ch <- ev:
		default:
		}
	}
}
