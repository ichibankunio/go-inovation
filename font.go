package main

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

type Font struct {
	fonts map[rune]*ebiten.Image
}

func NewFont() *Font {
	return &Font{
		fonts: map[rune]*ebiten.Image{},
	}
}

func (f *Font) DrawNumber(target *ebiten.Image, num int, x, y int) error {
	msg := strconv.Itoa(num)
	for _, c := range msg {
		if c == 32 {
			x += 9
			continue
		}
		if img, ok := f.fonts[c]; ok {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			if err := target.DrawImage(img, op); err != nil {
				return err
			}
			w, _ := img.Size()
			x += w
			continue
		}
		panic(fmt.Sprintf("DrawNumber couldn't find font file (%d) for number %d", c, num))
	}
	return nil
}
