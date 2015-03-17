package rkit

import (
	"github.com/gopherjs/gopherjs/js"
)

func init() {

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
