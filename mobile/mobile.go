package mobile

import (
	"github.com/hajimehoshi/ebiten/mobile"

	"github.com/hajimehoshi/go-inovation/ino"
)

const (
	ScreenWidth  = ino.ScreenWidth
	ScreenHeight = ino.ScreenHeight
)

func init() {
	game, err := ino.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetUpdateFunc(game.Loop)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
