package main

import (
	"fmt"

	"github.com/amitu/rkit"
)

func main() {
	data, err := rkit.LoadFile("words.gz")
	fmt.Println(err, len(data))
}
