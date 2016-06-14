// +build android

package inovation

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/hajimehoshi/go-inovation/ino"
)

type EventDispatcher mobile.EventDispatcher

const (
	ScreenWidth = ino.ScreenWidth
	ScreenHeight = ino.ScreenHeight
)

var (
	running         bool
	eventDispatcher EventDispatcher
)

func IsRunning() bool {
	return running
}

func Start(scale int) error {
	running = true
	game, err := ino.NewGame()
	if err != nil {
		return err
	}
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
