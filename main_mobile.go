// +build android

package inovation

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/hajimehoshi/go-inovation/ino"
)

type EventDispatcher mobile.EventDispatcher

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Start(width, height int) (EventDispatcher, error) {
	game, err := ino.NewGame()
	if err != nil {
		return nil, err
	}
	scale := max(1, min(width/ino.ScreenWidth, height/ino.ScreenHeight))
	e, err := mobile.Start(game.Loop, ino.ScreenWidth, ino.ScreenHeight, scale, ino.Title)
	if err != nil {
		return nil, err
	}
	return e, nil
}
