package draw

import (
	"bytes"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/go-inovation/ino/internal/assets"
	"github.com/hajimehoshi/go-inovation/ino/internal/input"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
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

var (
	images = map[string]*ebiten.Image{}
)

func LoadImages() error {
	for _, f := range []string{"ino", "msg", "bg", "touch"} {
		b, err := assets.Asset("resources/images/color/" + f + ".png")
		if err != nil {
			return err
		}
		origImg, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			return err
		}
		images[f], _ = ebiten.NewImageFromImage(origImg, ebiten.FilterDefault)
	}
	return nil
}

func Draw(screen *ebiten.Image, key string, px, py, sx, sy, sw, sh int) {
	op := &ebiten.DrawImageOptions{}
	r := image.Rect(sx, sy, sx+sw, sy+sh)
	op.SourceRect = &r
	op.GeoM.Translate(float64(px), float64(py))
	screen.DrawImage(images[key], op)
}

func DrawTouchButtons(screen *ebiten.Image) {
	img := images["touch"]
	w, h := img.Size()
	w /= 4
	dx := 0
	dy := ScreenHeight - h
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.4)
	for _, i := range []int{0, 1, 3} {
		r := image.Rect(i*w, 0, (i+1)*w, h)
		op.SourceRect = &r
		op.GeoM.Reset()
		op.GeoM.Translate(float64(dx+i*w), float64(dy))
		screen.DrawImage(img, op)
	}
	// Render 'down' button
	op = &ebiten.DrawImageOptions{}
	r := image.Rect(2*w, 0, 3*w, h)
	op.SourceRect = &r
	op.GeoM.Translate(float64(dx+2*w), float64(dy))
	alpha := 0.0
	if input.Current().IsActionKeyPressed() {
		alpha = 0.4
	} else {
		alpha = 0.1
	}
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(img, op)
}
