package mobile

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/mobile"

	"github.com/hajimehoshi/go-inovation/ino"
)

type game struct {
	g *ino.Game
}

func (g *game) Update(screen *ebiten.Image) error {
	return g.g.Loop(screen)
}

func (g *game) Layout(viewWidth, viewHeight int) (screenWidth, screenHeight int) {
	return ino.ScreenWidth, ino.ScreenHeight
}

func init() {
	inogame, err := ino.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(&game{g: inogame})
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
