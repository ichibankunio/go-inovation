package ino

import (
	"github.com/hajimehoshi/ebiten"
)

var input = &Input{
	pressed:     map[ebiten.Key]struct{}{},
	prevPressed: map[ebiten.Key]struct{}{},
}

var keys = []ebiten.Key{
	ebiten.KeyEnter,
	ebiten.KeySpace,
	ebiten.KeyLeft,
	ebiten.KeyDown,
	ebiten.KeyRight,
}

type Input struct {
	pressed      map[ebiten.Key]struct{}
	prevPressed  map[ebiten.Key]struct{}
	touchEnabled bool
}

func (i *Input) TouchEnabled() bool {
	if isTouchPrimaryInput() {
		return true
	}
	return i.touchEnabled
}

func (i *Input) Update() {
	i.prevPressed = map[ebiten.Key]struct{}{}
	for k, _ := range i.pressed {
		i.prevPressed[k] = struct{}{}
	}
	i.pressed = map[ebiten.Key]struct{}{}
	for _, k := range keys {
		if ebiten.IsKeyPressed(k) {
			i.pressed[k] = struct{}{}
		}
	}
	// Emulates the keys by touching
	touches := ebiten.Touches()
	for _, t := range touches {
		x, _ := t.Position()
		switch {
		case 320 <= x:
		case 240 <= x:
			i.pressed[ebiten.KeyEnter] = struct{}{}
			i.pressed[ebiten.KeySpace] = struct{}{}
		case 160 <= x:
			i.pressed[ebiten.KeyDown] = struct{}{}
		case 80 <= x:
			i.pressed[ebiten.KeyRight] = struct{}{}
		default:
			i.pressed[ebiten.KeyLeft] = struct{}{}
		}
	}
	if 0 < len(touches) {
		i.touchEnabled = true
	}
}

func (i *Input) IsKeyPressed(key ebiten.Key) bool {
	_, ok := i.pressed[key]
	return ok
}

// TODO(hajimehoshi): Rename this to IsKeyTrigger?
func (i *Input) IsKeyPushed(key ebiten.Key) bool {
	_, ok := i.pressed[key]
	if !ok {
		return false
	}
	_, ok = i.prevPressed[key]
	return !ok
}

func (i *Input) IsActionKeyPressed() bool {
	return i.IsKeyPressed(ebiten.KeyEnter) || i.IsKeyPressed(ebiten.KeySpace)
}

func (i *Input) IsActionKeyPushed() bool {
	return i.IsKeyPushed(ebiten.KeyEnter) || i.IsKeyPushed(ebiten.KeySpace)
}
