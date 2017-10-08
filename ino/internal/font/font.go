package font

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

type Font struct {
	imgs map[rune]*ebiten.Image
}

func NewFont(imgs map[rune]*ebiten.Image) *Font {
	return &Font{
		imgs: imgs,
	}
}

func (f *Font) DrawNumber(target *ebiten.Image, num int, x, y int) {
	msg := strconv.Itoa(num)
	for _, c := range msg {
		if c == 32 {
			x += 9
			continue
		}
		if img, ok := f.imgs[c]; ok {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			target.DrawImage(img, op)
			w, _ := img.Size()
			x += w
			continue
		}
		panic(fmt.Sprintf("DrawNumber couldn't find font file (%d) for number %d", c, num))
	}
}
