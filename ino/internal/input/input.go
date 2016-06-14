package input

import (
	"github.com/hajimehoshi/ebiten"
)

var theInput = &Input{
	pressed:     map[ebiten.Key]struct{}{},
	prevPressed: map[ebiten.Key]struct{}{},
}

func Current() *Input {
	return theInput
}

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionDown
)

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

func (i *Input) IsTouchEnabled() bool {
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
		x, y := t.Position()
		// TODO(hajimehoshi): 240 and 64 are magic numbers
		if y < 240-64 {
			continue
		}
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

func (i *Input) IsSpaceTouched() bool {
	for _, t := range ebiten.Touches() {
		_, y := t.Position()
		if y < 240-64 {
			return true
		}
	}
	return false
}

func (i *Input) isKeyPressed(key ebiten.Key) bool {
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
	return i.isKeyPressed(ebiten.KeyEnter) || i.isKeyPressed(ebiten.KeySpace)
}

func (i *Input) IsActionKeyPushed() bool {
	return i.IsKeyPushed(ebiten.KeyEnter) || i.IsKeyPushed(ebiten.KeySpace)
}

func (i *Input) IsDirectionKeyPressed(dir Direction) bool {
	switch dir {
	case DirectionLeft:
		return i.isKeyPressed(ebiten.KeyLeft)
	case DirectionRight:
		return i.isKeyPressed(ebiten.KeyRight)
	case DirectionDown:
		return i.isKeyPressed(ebiten.KeyDown)
	default:
		panic("not reach")
	}
}
