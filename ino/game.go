package ino

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/go-inovation/ino/internal/assets"
	"github.com/hajimehoshi/go-inovation/ino/internal/input"
)

var (
	imageItemFrame *ebiten.Image
)

func init() {
	imageItemFrame, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
	imageItemFrame.Fill(color.Black)
	ebitenutil.DrawRect(imageItemFrame, 2, 2, 28, 28, color.White)
}

func (g *Game) DrawItemFrame(x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	g.screen.DrawImage(imageItemFrame, op)
}

type imgPart struct {
	px, py, sx, sy, sw, sh int
}

func (g *Game) Draw(key string, px, py, sx, sy, sw, sh int) {
	op := &ebiten.DrawImageOptions{}
	r := image.Rect(sx, sy, sx+sw, sy+sh)
	op.SourceRect = &r
	op.GeoM.Translate(float64(px), float64(py))
	g.screen.DrawImage(g.img[key], op)
}

func (g *Game) DrawParts(key string, parts []imgPart) {
	op := &ebiten.DrawImageOptions{}
	for _, p := range parts {
		r := image.Rect(p.sx, p.sy, p.sx+p.sw, p.sy+p.sh)
		op.SourceRect = &r
		op.GeoM.Reset()
		op.GeoM.Translate(float64(p.px), float64(p.py))
		g.screen.DrawImage(g.img[key], op)
	}
}

func (g *Game) DrawNumber(num int, x, y int) {
	g.font.DrawNumber(g.screen, num, x, y)
}

func (g *Game) DrawTouchButtons() {
	img := g.img["touch"]
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
		g.screen.DrawImage(img, op)
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
	g.screen.DrawImage(img, op)
}

func (g *Game) loadImages() error {
	for _, f := range []string{"ino", "msg", "bg", "touch"} {
		b, err := assets.Asset("resources/images/color/" + f + ".png")
		if err != nil {
			return err
		}
		origImg, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			return err
		}
		img, err := ebiten.NewImageFromImage(origImg, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		g.img[f] = img
	}

	g.font = NewFont()
	for n := 48; n <= 57; n++ {
		b, err := assets.Asset(fmt.Sprintf("resources/font/%d.png", n))
		if err != nil {
			return err
		}
		origImg, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			return err
		}
		img, err := ebiten.NewImageFromImage(origImg, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		g.font.fonts[rune(n)] = img
	}
	return nil
}

func NewGame() (game *Game, err error) {
	defer func() {
		if ferr := finalizeAudio(); ferr != nil && err == nil {
			err = ferr
		}
		if err != nil {
			game = nil
		}
	}()
	game = &Game{
		img:           map[string]*ebiten.Image{},
		imageLoadedCh: make(chan error),
		audioLoadedCh: make(chan error),
	}
	go func() {
		if err := game.loadImages(); err != nil {
			game.imageLoadedCh <- err
		}
		close(game.imageLoadedCh)
	}()
	go func() {
		if err := loadAudio(); err != nil {
			game.audioLoadedCh <- err
		}
		close(game.audioLoadedCh)
	}()
	return
}
