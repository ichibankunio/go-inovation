package font

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/go-mplusbitmap"
)

func DrawText(target *ebiten.Image, str string, x, y int) {
	// Adjust position for 'dot' position.
	x += 3
	y += 12
	text.Draw(target, str, mplusbitmap.Gothic12r, x, y, color.Black)
}
