package main

import (
	"fmt"
	"time"

	"github.com/amitu/rkit"
	"github.com/amitu/rkit/cursor"
)

func main() {
	rkit.SetCursor(cursor.Cell)
	for {
		fmt.Println(rkit.GetCursor().String())
		time.Sleep(time.Second)
		rkit.SetCursor(cursor.NResize)
		time.Sleep(time.Second)
		rkit.SetCursor(cursor.SResize)
	}
}
