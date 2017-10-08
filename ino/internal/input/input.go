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

	// Fullscreen
	ebiten.KeyF,

	// Profiling
	ebiten.KeyP,
	ebiten.KeyQ,
}

type Input struct {
	pressed          map[ebiten.Key]struct{}
	prevPressed      map[ebiten.Key]struct{}
	spaceTouched     bool
	prevSpaceTouched bool
	touchEnabled     bool
}

func (i *Input) IsTouchEnabled() bool {
	if isTouchPrimaryInput() {
		return true
	}
	return i.touchEnabled
}

func (i *Input) Update() {
	const gamepadID = 0

	i.prevPressed = map[ebiten.Key]struct{}{}
	for k := range i.pressed {
		i.prevPressed[k] = struct{}{}
	}
	i.pressed = map[ebiten.Key]struct{}{}
	for _, k := range keys {
		if ebiten.IsKeyPressed(k) {
			i.pressed[k] = struct{}{}
		}
	}

	// Emulates the keys by gamepad pressing
	switch ebiten.GamepadAxis(gamepadID, 0) {
	case -1:
		i.pressed[ebiten.KeyLeft] = struct{}{}
	case 1:
		i.pressed[ebiten.KeyRight] = struct{}{}
	}
	if y := ebiten.GamepadAxis(gamepadID, 1); y == 1 {
		i.pressed[ebiten.KeyDown] = struct{}{}
	}
	for b := ebiten.GamepadButton0; b <= ebiten.GamepadButtonMax; b++ {
		if ebiten.IsGamepadButtonPressed(gamepadID, b) {
			i.pressed[ebiten.KeyEnter] = struct{}{}
			i.pressed[ebiten.KeySpace] = struct{}{}
			break
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
	i.prevSpaceTouched = i.spaceTouched
	i.spaceTouched = false
	for _, t := range touches {
		_, y := t.Position()
		if y < 240-64 {
			i.spaceTouched = true
			break
		}
	}
	if 0 < len(touches) {
		i.touchEnabled = true
	}
}

func (i *Input) IsSpaceTouched() bool {
	return i.spaceTouched
}

func (i *Input) IsSpaceJustTouched() bool {
	return i.spaceTouched && !i.prevSpaceTouched
}

func (i *Input) IsKeyPressed(key ebiten.Key) bool {
	_, ok := i.pressed[key]
	return ok
}

func (i *Input) IsKeyJustPressed(key ebiten.Key) bool {
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

func (i *Input) IsActionKeyJustPressed() bool {
	return i.IsKeyJustPressed(ebiten.KeyEnter) || i.IsKeyJustPressed(ebiten.KeySpace)
}

func (i *Input) IsDirectionKeyPressed(dir Direction) bool {
	switch dir {
	case DirectionLeft:
		return i.IsKeyPressed(ebiten.KeyLeft)
	case DirectionRight:
		return i.IsKeyPressed(ebiten.KeyRight)
	case DirectionDown:
		return i.IsKeyPressed(ebiten.KeyDown)
	default:
		panic("not reach")
	}
}
