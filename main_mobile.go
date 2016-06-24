// +build android ios darwin,arm darwin,arm64

package inovation

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/hajimehoshi/go-inovation/ino"
)

const (
	ScreenWidth  = ino.ScreenWidth
	ScreenHeight = ino.ScreenHeight
)

var (
	running         bool
	eventDispatcher mobile.EventDispatcher
)

func IsRunning() bool {
	return running
}

func Start(scale float64) error {
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

func Render() error {
	return eventDispatcher.Render()
}

func UpdateTouchesOnAndroid(action int, id int, x, y int) {
	eventDispatcher.UpdateTouchesOnAndroid(action, id, x, y)
}

func UpdateTouchesOnIOS(phase int, ptr int, x, y int) {
	eventDispatcher.UpdateTouchesOnIOS(phase, ptr, x, y)
}
