package main

import (
	"fmt"
	"time"

	"github.com/amitu/rkit"
)

func main() {
	rkit.SetTitle("yo")

	go func() {
		for {
			<-rkit.DesktopResize
			fmt.Println("got \"desktop\" resize evt")
		}
	}()

	// go func() {
	// 	c := rkit.DesktopResize.Sub()
	// 	for {
	// 		<-c
	// 		fmt.Println("got evt")
	// 	}
	// 	rkit.DesktopResize.Unsub(c)
	// }()

	// go func() {
	// 	c := rkit.DesktopResize.Sub()
	// 	for {
	// 		<-c
	// 		fmt.Println("got evt")
	// 	}
	// 	rkit.DesktopResize.Unsub(c)
	// }()

	for {
		fmt.Println(rkit.Width(), rkit.Height(), rkit.Title())
		rkit.SetTitle(fmt.Sprintf("%d", time.Now().UnixNano()))
		time.Sleep(time.Second)
	}
}
