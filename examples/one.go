package main

import (
	"fmt"
	"time"

	"github.com/amitu/rkit"
)

func main() {
	rkit.SetTitle("yo")

	go func() {
		ch := rkit.Mouse.Sub()
		for {
			ev := <-ch
			m := ev.(rkit.MouseEvent)
			fmt.Printf("got mouse: %d, %d, %d\n", m.X, m.Y, m.Action)
		}
	}()

	go func() {
		ch := rkit.Key.Sub()
		for {
			ev := <-ch
			key := ev.(rkit.KeyEvent)
			fmt.Printf("got key: %#U, %d, %d\n", key.Char, key.Code, key.Action)
		}
	}()

	go func() {
		ch := rkit.DesktopResize.Sub()
		for {
			<-ch
			fmt.Println("got evt")
		}
	}()

	go func() {
		ch := rkit.DesktopResize.Sub()
		for {
			<-ch
			fmt.Println("got evt22")
		}
	}()

	// go func() {
	// 	ch := rkit.AnimationFrame.Sub()
	// 	for {
	// 		<-ch
	// 		fmt.Println("got frame")
	// 	}
	// }()

	// go func() {
	// 	ch := rkit.AnimationFrame.Sub()
	// 	for {
	// 		<-ch
	// 		fmt.Println("got frame2")
	// 	}
	// }()

	for {
		// fmt.Println(rkit.Width(), rkit.Height(), rkit.Title())
		rkit.SetTitle(fmt.Sprintf("%d", time.Now().UnixNano()))
		time.Sleep(time.Second)
	}
}
