package rkit

// https://github.com/ajhager/raf/blob/master/raf.go

import (
	"fmt"

	"github.com/ajhager/raf"
)

var (
	AnimationFrame *EventSource
	keepWatching   = false
)

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
