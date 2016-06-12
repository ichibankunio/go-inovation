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

var (
	running         bool
	eventDispatcher EventDispatcher
)

func IsRunning() bool {
	return running
}

func Start(width, height int) error {
	running = true
	game, err := ino.NewGame()
	if err != nil {
		return err
	}
	scale := max(1, min(width/ino.ScreenWidth, height/ino.ScreenHeight))
	e, err := mobile.Start(game.Loop, ino.ScreenWidth, ino.ScreenHeight, scale, ino.Title)
	if err != nil {
		return err
	}
	eventDispatcher = e
	return nil
}

func CurrentEventDispatcher() EventDispatcher {
	return eventDispatcher;
}
