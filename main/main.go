package main

import (
	"github.com/hajimehoshi/inovation5"
)

func main() {
	if err := inovation5.Run(); err != nil {
		panic(err)
	}
}
