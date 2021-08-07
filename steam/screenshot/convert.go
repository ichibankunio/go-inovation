package main

import (
	"errors"
	"fmt"
	"image/png"
	"os"
	"regexp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var regularTermination = errors.New("regular termiination")

var reFilename = regexp.MustCompile(`^screenshot(.+)_([a-z]+)\.png$`)

type Game struct {
}

func (*Game) Update() error {
	ents, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	const (
		screenshotWidth  = 1920
		screenshotHeight = 1080
	)

	offscreen := ebiten.NewImage(screenshotWidth, screenshotHeight)
	for _, ent := range ents {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		m := reFilename.FindStringSubmatch(name)
		if m == nil {
			continue
		}

		img, _, err := ebitenutil.NewImageFromFile(name)
		if err != nil {
			return err
		}

		const scale = 6
		//w, h := img.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(0, -192)

		offscreen.Clear()
		offscreen.DrawImage(img, op)

		newName := fmt.Sprintf("screenshot%s_%dx%d_%s.png", m[0], screenshotWidth, screenshotHeight, m[1])
		f, err := os.Create(newName)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := png.Encode(f, offscreen); err != nil {
			return err
		}
	}
	return regularTermination
}

func (*Game) Draw(screen *ebiten.Image) {
}

func (*Game) Layout(width, height int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(&Game{}); err != nil && err != regularTermination {
		panic(err)
	}
}
