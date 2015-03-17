package main

import (
	"fmt"
	"time"

	"github.com/amitu/rkit"
)

func main() {
	rkit.SetTitle("yo")
	for {
		fmt.Println(rkit.Width(), rkit.Height(), rkit.Title())
		time.Sleep(time.Second)
	}
}
