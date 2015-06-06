package inovation5

import (
	"fmt"
	"path/filepath"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Font struct {
	fonts map[rune]*ebiten.Image
}

func NewFont() *Font {
	return &Font{
		fonts: map[rune]*ebiten.Image{},
	}
}

func (f *Font) Load(path string) error {
	// TODO(hajimehoshi): Use goroutine
	for n := 33; n < 126; n++ {
		src := filepath.Join(path, fmt.Sprintf("%d.png", n))
		img, _, err := ebitenutil.NewImageFromFile(src, ebiten.FilterNearest)
		if err != nil {
			return err
		}
		f.fonts[rune(n)] = img
	}
	return nil
}
