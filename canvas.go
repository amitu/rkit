package rkit

import (
	"github.com/amitu/rkit/cursor"
	"github.com/gopherjs/gopherjs/js"
)

var (
	// DesktopResize is an "EventSource" that can be used to subscribe to
	// "Desktop" resizing events.
	//
	// This is not the same as window resize, as on one desktop one may have
	// multiple windows. We do not even have a concept of window so far, but
	// when do, that would be different than this.
	DesktopResize *EventSource

	// Key is an EventSource, that can be used to subscribe to keyboard events.
	//
	// This is a top level event. In most cases you do not need this, and want a
	// widget specific EventSource for keys.
	Key *EventSource

	// Mouse is an EventSource, that can be used to subscribe to mous/touch
	// events.
	//
	// This is a top level event. In most cases you do not need this, and want a
	// widget specific EventSource for mouse/touch event.
	Mouse    *EventSource
	canvas   *js.Object
	textarea *js.Object

	cursorMap map[string]cursor.Cursor
)

// DesktopResizeEvent struct. This would be the concrete struct passed to
// subscribers of DesktopResize EventSource.
type DesktopResizeEvent struct {
	BaseEvent
}

// Action is available on Key and Mouse events as .Action. It can be used to
// determine the action for the event.
type Action int

var (
	// MOVE is Mouse or Touch move Action
	MOVE = Action(0)
	// PRESS is Mouse, Touch, or Keyboard press Action
	PRESS = Action(1)
	// RELEASE is Mouse, Touch or Keyboard release Action
	RELEASE = Action(2)
)

// KeyEvent is the struct that is sent over on Key event channel when a keyboard
// event occurs.
type KeyEvent struct {
	BaseEvent
	Char      rune
	Code      int
	Modifiers int
	Action    Action
}

// MouseEvent is the struct that is sent over on Mouse event channel when a
// mouse or touch event occurs.
type MouseEvent struct {
	BaseEvent
	X, Y   int
	Action Action
}

func init() {
	DesktopResize = MakeEventSource()
	Key = MakeEventSource()
	Mouse = MakeEventSource()

	initCanvas()
	initEvents()

	cursorMap = map[string]cursor.Cursor{
		"alias":         cursor.Alias,
		"all-scroll":    cursor.AllScroll,
		"auto":          cursor.Auto,
		"cell":          cursor.Cell,
		"context-menu":  cursor.ContextMenu,
		"col-resize":    cursor.ColResize,
		"copy":          cursor.Copy,
		"crosshair":     cursor.CrossHair,
		"default":       cursor.Default,
		"e-resize":      cursor.EResize,
		"ew-resize":     cursor.EWResize,
		"grab":          cursor.Grab,
		"grabbing":      cursor.Grabbing,
		"help":          cursor.Help,
		"move":          cursor.Move,
		"n-resize":      cursor.NResize,
		"ne-resize":     cursor.NEResize,
		"nesw-resize":   cursor.NESWResize,
		"ns-resize":     cursor.NSResize,
		"nw-resize":     cursor.NWResize,
		"nwse-resize":   cursor.NWSEResize,
		"no-drop":       cursor.NoDrop,
		"none":          cursor.None,
		"not-allowed":   cursor.NotAllowed,
		"pointer":       cursor.Pointer,
		"progress":      cursor.Progress,
		"row-resize":    cursor.RowResize,
		"s-resize":      cursor.SResize,
		"se-resize":     cursor.SEResize,
		"sw-resize":     cursor.SWResize,
		"text":          cursor.Text,
		"vertical-text": cursor.VerticalText,
		"w-resize":      cursor.WResize,
		"wait":          cursor.Wait,
		"zoom-in":       cursor.ZoomIn,
		"zoom-out":      cursor.ZoomOut,
		"initial":       cursor.Initial,
		"inherit":       cursor.Inherit,
	}
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

	// mouse events
	textarea.Call("addEventListener", "mousemove", func(ev *js.Object) {
		rect := textarea.Call("getBoundingClientRect")
		Mouse.Pub(
			MouseEvent{
				X:      ev.Get("clientX").Int() - rect.Get("left").Int(),
				Y:      ev.Get("clientY").Int() - rect.Get("top").Int(),
				Action: MOVE,
			},
		)
	}, false)

	textarea.Call("addEventListener", "mousedown", func(ev *js.Object) {
		rect := textarea.Call("getBoundingClientRect")
		Mouse.Pub(
			MouseEvent{
				X:      ev.Get("clientX").Int() - rect.Get("left").Int(),
				Y:      ev.Get("clientY").Int() - rect.Get("top").Int(),
				Action: PRESS,
			},
		)
	}, false)

	textarea.Call("addEventListener", "mouseup", func(ev *js.Object) {
		rect := textarea.Call("getBoundingClientRect")
		Mouse.Pub(
			MouseEvent{
				X:      ev.Get("clientX").Int() - rect.Get("left").Int(),
				Y:      ev.Get("clientY").Int() - rect.Get("top").Int(),
				Action: RELEASE,
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

// Width returns the width of the "desktop". Note this is not same as window.
func Width() int {
	return js.Global.Get("window").Get("innerWidth").Int()
}

// Height returns the height of the "desktop". Note this is not same as window.
func Height() int {
	return js.Global.Get("window").Get("innerHeight").Int()
}

// Title returns the title of the "desktop". Note this is not same as window.
// This may not be supported on all platforms.
func Title() string {
	return js.Global.Get("document").Get("title").String()
}

// SetTitle changes the title "desktop". This may not be supported in all
// platforms. In case of browswer, the title is shown in tab, but in iphone etc
// it may not be shown anywhere.
func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}

func cursor2css(c cursor.Cursor) string {
	for css, cur := range cursorMap {
		if cur == c {
			return css
		}
	}
	return ""
}

// GetCursor returns a cursor.Cursor object returning the current cursor.
func GetCursor() cursor.Cursor {
	tstyle := textarea.Get("style")
	return cursorMap[tstyle.Get("cursor").String()]
}

// SetCursor takes a cursor.Cursor and makes it current.
func SetCursor(c cursor.Cursor) {
	tstyle := textarea.Get("style")
	tstyle.Set("cursor", cursor2css(c))
}
