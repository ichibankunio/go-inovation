// +build js

package ino

import (
	"github.com/hajimehoshi/ebiten"
)

func Scale() float64 {
	return 1
}

func init() {
	ebiten.SetFullscreen(true)
}
