package mobile

import (
	"github.com/hajimehoshi/ebiten/mobile"

	"github.com/hajimehoshi/go-inovation/ino"
)

const (
	ScreenWidth  = ino.ScreenWidth
	ScreenHeight = ino.ScreenHeight
)

func Start(scale float64) error {
	game, err := ino.NewGame()
	if err != nil {
		return err
	}
	if err := mobile.Start(game.Loop, ino.ScreenWidth, ino.ScreenHeight, scale, ino.Title); err != nil {
		return err
	}
	return nil
}

func Update() error {
	return mobile.Update()
}

func UpdateTouchesOnAndroid(action int, id int, x, y int) {
	mobile.UpdateTouchesOnAndroid(action, id, x, y)
}

func UpdateTouchesOnIOS(phase int, ptr int64, x, y int) {
	mobile.UpdateTouchesOnIOS(phase, ptr, x, y)
}
