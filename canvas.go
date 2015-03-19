package rkit

import "github.com/gopherjs/gopherjs/js"

var (
	/*
		DesktopResize is an "EventSource" that can be used to subscribe to
		"Desktop" resizing events.

		This is not the same as window resize, as on one desktop one may have
		multiple windows. We do not even have a concept of window so far, but
		when do, that would be different than this.
	*/
	DesktopResize *EventSource
	canvas        *js.Object
)

/*
	DesktopResizeEvent struct. This would be the concrete struct passed to
	subscribers of DesktopResize EventSource.
*/
type DesktopResizeEvent struct {
	BaseEvent
}

func init() {
	DesktopResize = MakeEventSource()

	js.Global.Get("window").Call(
		"addEventListener", "resize", func() {
			DesktopResize.Pub(DesktopResizeEvent{})
		},
	)

	initCanvas()
}

func initCanvas() {
	document := js.Global.Get("document")

	body := document.Get("body")
	bstyle := body.Get("style")
	bstyle.Set("display", "block")
	bstyle.Set("width", "100%")
	bstyle.Set("height", "100%")
	bstyle.Set("margin", "0px")
	bstyle.Set("padding", "0px")

	canvas = document.Call("createElement", "canvas")
	cstyle := canvas.Get("style")
	cstyle.Set("display", "block")
	cstyle.Set("width", "100%")
	cstyle.Set("height", "100%")
	cstyle.Set("margin", "0px")
	cstyle.Set("padding", "0px")

	body.Call("appendChild", canvas)
}

/*
	Width() returns the width of the "desktop". Note this is not same as window.
*/
func Width() int {
	return js.Global.Get("window").Get("innerWidth").Int()
}

/*
	Height() returns the height of the "desktop". Note this is not same as
	window.
*/
func Height() int {
	return js.Global.Get("window").Get("innerHeight").Int()
}

/*
	Title() returns the title of the "desktop". Note this is not same as window.
	This may not be supported on all platforms.
*/
func Title() string {
	return js.Global.Get("document").Get("title").String()
}

/*
	SetTitle() changes the title "desktop". This may not be supported in all
	platforms. In case of browswer, the title is shown in tab, but in iphone etc
	it may not be shown anywhere.
*/
func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}
