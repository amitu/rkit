package rkit

import "github.com/gopherjs/gopherjs/js"

var (
	DesktopResize chan struct{}
)

func init() {
	DesktopResize = make(chan struct{})

	js.Global.Get("window").Call(
		"addEventListener", "resize", func() {
			select {
			case DesktopResize <- struct{}{}:
			default:
				break
			}
		},
	)
}

func Width() int {
	return js.Global.Get("window").Get("innerWidth").Int()
}

func Height() int {
	return js.Global.Get("window").Get("innerHeight").Int()
}

func Title() string {
	return js.Global.Get("document").Get("title").String()
}

func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}
