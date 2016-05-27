package ino

import (
	"github.com/hajimehoshi/ebiten"
)

var input = &Input{}

// TODO(hajimehoshi): 256 is an arbitrary number.
const maxKey = 256

type Input struct {
	pressed     [maxKey]bool
	prevPressed [maxKey]bool
}

func (i *Input) Update() {
	i.prevPressed = i.pressed
	for k, _ := range i.pressed {
		k := ebiten.Key(k)
		i.pressed[k] = ebiten.IsKeyPressed(k)
	}
}

func (i *Input) IsKeyPressed(key ebiten.Key) bool {
	return i.pressed[key]
}

// TODO(hajimehoshi): Rename this to IsKeyTrigger?
func (i *Input) IsKeyPushed(key ebiten.Key) bool {
	return i.pressed[key] && !i.prevPressed[key]
}

func (i *Input) IsActionKeyPressed() bool {
	return i.IsKeyPressed(ebiten.KeyEnter) || i.IsKeyPressed(ebiten.KeySpace)
}

func (i *Input) IsActionKeyPushed() bool {
	return i.IsKeyPushed(ebiten.KeyEnter) || i.IsKeyPushed(ebiten.KeySpace)
}
