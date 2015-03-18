package main

import (
	"fmt"
	"time"

	"github.com/amitu/rkit"
)

func main() {
	rkit.SetTitle("yo")

	go func() {
		ch := rkit.DesktopResize.Sub()
		for {
			<-ch
			fmt.Println("got evt")
		}
		rkit.DesktopResize.Unsub(ch)
	}()

	go func() {
		ch := rkit.DesktopResize.Sub()
		for {
			<-ch
			fmt.Println("got evt22")
		}
		rkit.DesktopResize.Unsub(ch)
	}()

	go func() {
		ch := rkit.AnimationFrame.Sub()
		for {
			<-ch
			fmt.Println("got frame")
		}
		rkit.DesktopResize.Unsub(ch)
	}()

	go func() {
		ch := rkit.AnimationFrame.Sub()
		for {
			<-ch
			fmt.Println("got frame")
		}
		rkit.DesktopResize.Unsub(ch)
	}()

	for {
		fmt.Println(rkit.Width(), rkit.Height(), rkit.Title())
		rkit.SetTitle(fmt.Sprintf("%d", time.Now().UnixNano()))
		time.Sleep(time.Second)
	}
}
