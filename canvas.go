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
	/*
		Key is an EventSource, that can be used to subscribe to keyboard events.

		This is a top level event. In most cases you do not need this, and want
		a widget specific EventSource for keys.
	*/
	Key *EventSource
	/*
		Mouse is an EventSource, that can be used to subscribe to mous/touch
		events.

		This is a top level event. In most cases you do not need this, and want
		a widget specific EventSource for mouse/touch event.
	*/
	Mouse    *EventSource
	canvas   *js.Object
	textarea *js.Object
)

/*
	DesktopResizeEvent struct. This would be the concrete struct passed to
	subscribers of DesktopResize EventSource.
*/
type DesktopResizeEvent struct {
	BaseEvent
}

/*
	Key and Mouse events have an attribute .Action of type Action. It can be
	used to determine the action for the event.
*/
type Action int

var (
	MOVE    = Action(0)
	PRESS   = Action(1)
	RELEASE = Action(2)
)

/*
	KeyEvent is the struct that is sent over on Key event channel when a
	keyboard event occurs.
*/
type KeyEvent struct {
	BaseEvent
	Char      rune
	Code      int
	Modifiers int
	Action    Action
}

func init() {
	DesktopResize = MakeEventSource()
	Key = MakeEventSource()
	Mouse = MakeEventSource()

	initCanvas()
	initEvents()
}

func initEvents() {
	js.Global.Get("window").Call(
		"addEventListener", "resize", func() {
			DesktopResize.Pub(DesktopResizeEvent{})
		},
	)

	// http://unixpapa.com/js/key.html
	textarea.Call("addEventListener", "keydown", func(ev *js.Object) {
		Key.Pub(
			KeyEvent{
				Code:   ev.Get("keyCode").Int(),
				Char:   rune(ev.Get("charCode").Int()),
				Action: PRESS,
			},
		)
	}, false)

	textarea.Call("addEventListener", "keyup", func(ev *js.Object) {
		Key.Pub(
			KeyEvent{
				Code:   ev.Get("keyCode").Int(),
				Char:   rune(ev.Get("charCode").Int()),
				Action: RELEASE,
			},
		)
	}, false)

	textarea.Call("addEventListener", "keypress", func(ev *js.Object) {
		Key.Pub(
			KeyEvent{
				Code:   ev.Get("keyCode").Int(),
				Char:   rune(ev.Get("charCode").Int()),
				Action: PRESS,
			},
		)
	}, false)
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

	textarea = document.Call("createElement", "textarea")
	tstyle := textarea.Get("style")
	tstyle.Set("position", "absolute")
	tstyle.Set("top", "0")
	tstyle.Set("left", "0")
	tstyle.Set("display", "block")
	tstyle.Set("width", "100%")
	tstyle.Set("height", "100%")
	tstyle.Set("margin", "0px")
	tstyle.Set("padding", "0px")
	tstyle.Set("z-index", "1000")
	tstyle.Set("opacity", "0.0")
	tstyle.Set("filter", "alpha(opacity=0)") /* For IE8 and earlier */

	body.Call("appendChild", textarea)
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
