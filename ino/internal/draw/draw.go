package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	imageItemFrame *ebiten.Image
)

func init() {
	imageItemFrame, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
	imageItemFrame.Fill(color.Black)
	ebitenutil.DrawRect(imageItemFrame, 2, 2, 28, 28, color.White)
}

func DrawItemFrame(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(imageItemFrame, op)
}
