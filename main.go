package main

import (
	"github.com/hajimehoshi/go-inovation/ino"
)

func main() {
	if err := ino.Run(); err != nil {
		panic(err)
	}
}
