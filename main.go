// +build !android

package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/go-inovation/ino"
)

func main() {
	game, err := ino.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.Run(game.Loop, ino.ScreenWidth, ino.ScreenHeight, 2, ino.Title); err != nil {
		panic(err)
	}
}
