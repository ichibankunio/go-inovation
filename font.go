package main

import (
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
