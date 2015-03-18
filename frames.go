package rkit

// https://github.com/ajhager/raf/blob/master/raf.go

import (
	"fmt"

	"github.com/ajhager/raf"
)

var (
	/*
		AnimationFrame is an "EventSource" that can be used to subscribe to
		"EventFrame" events.

		EventFrame is an event that is fired when the underlying graphic library
		requests a new frame to be rendered.

		If you want to do any kind of animation, instead of running a go routine
		powered by a timer, power it by this event source. The timer you have
		may fire more often than the underlying system updates the screen, so it
		would be wastage of computational resources.
	*/
	AnimationFrame *EventSource
	keepWatching   = false
)

/*
	AnimationFrameEvent struct. This would be the concrete struct passed to
	subscribers of AnimationFrame EventSource.
*/
type AnimationFrameEvent struct {
	BaseEvent
}

func rafCallback(f float32) {
	fmt.Println("here")
	AnimationFrame.Lock.RLock()
	defer AnimationFrame.Lock.RUnlock()

	fmt.Println("here2")

	if keepWatching {
		fmt.Println("rafCa", f)
		AnimationFrame.Pub(AnimationFrameEvent{})
		raf.RequestAnimationFrame(rafCallback)
	}
}

type animationFrameWatcher struct {
}

func (a animationFrameWatcher) Start() {
	fmt.Println("AnimationFrameWatcher.Start()")
	keepWatching = true
	go rafCallback(0)
}

func (a animationFrameWatcher) Stop() {
	fmt.Println("AnimationFrameWatcher.Stop()")
	keepWatching = false
}

func init() {
	AnimationFrame = MakeEventSource()
	AnimationFrame.Watcher = &animationFrameWatcher{}
}
