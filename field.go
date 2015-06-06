package inovation5

import (
	"strings"
)

type FieldType int

const (
	FIELD_NONE         FieldType = iota // なし
	FIELD_HIDEPATH                      // 隠しルート(見えるけど判定のないブロック)
	FIELD_UNVISIBLE                     // 不可視ブロック(見えないけど判定があるブロック)
	FIELD_BLOCK                         // 通常ブロック
	FIELD_BAR                           // 床。降りたり上ったりできる
	FIELD_SCROLL_L                      // ベルト床左
	FIELD_SCROLL_R                      // ベルト床右
	FIELD_SPIKE                         // トゲ
	FIELD_SLIP                          // すべる
	FIELD_ITEM_BORDER                   // アイテムチェック用
	FIELD_ITEM_POWERUP                  // パワーアップ
	// ふじ系
	FIELD_ITEM_FUJI
	FIELD_ITEM_BUSHI
	FIELD_ITEM_APPLE
	FIELD_ITEM_V
	// たか系
	FIELD_ITEM_TAKA
	FIELD_ITEM_SHUOLDER
	FIELD_ITEM_DAGGER
	FIELD_ITEM_KATAKATA
	// なす系
	FIELD_ITEM_NASU
	FIELD_ITEM_BONUS
	FIELD_ITEM_NURSE
	FIELD_ITEM_NAZUNA
	// くそげー系
	FIELD_ITEM_GAMEHELL
	FIELD_ITEM_GUNDAM
	FIELD_ITEM_POED
	FIELD_ITEM_MILESTONE
	FIELD_ITEM_1YEN
	FIELD_ITEM_TRIANGLE
	FIELD_ITEM_OMEGA      // 隠し
	FIELD_ITEM_LIFE       // ハート
	FIELD_ITEM_STARTPOINT // 開始地点
	FIELD_ITEM_MAX
)

const (
	FIELD_X_MAX       = 128
	FIELD_Y_MAX       = 128
	GRAPHIC_OFFSET_X  = -16 - 16*2
	GRAPHIC_OFFSET_Y  = 8 - 16*2
	SCROLLPANEL_SPEED = 2.0
)

type Position struct {
	X int
	Y int
}

type PositionF struct {
	X float64
	Y float64
}

type Field struct {
	field [FIELD_X_MAX * FIELD_Y_MAX]FieldType
	timer int
}

func NewField(data string) *Field {
	f := &Field{}
	xm := strings.Split(data, "\n")
	const decoder = " HUB~<>*I PabcdefghijklmnopqrzL@"

	for yy, line := range xm {
		for xx, c := range line {
			n := strings.IndexByte(decoder, byte(c))
			f.field[yy*FIELD_X_MAX+xx] = FieldType(n)
		}
	}
	return f
}

func (f *Field) Move() {
	f.timer++
}

func (f *Field) GetStartPoint() Position {
	p := Position{}
	for yy := 0; yy < FIELD_Y_MAX; yy++ {
		for xx := 0; xx < FIELD_X_MAX; xx++ {
			if f.GetField(xx, yy) == FIELD_ITEM_STARTPOINT {
				p.X = xx * CHAR_SIZE
				p.Y = yy * CHAR_SIZE
				f.EraseField(xx, yy)
				return p
			}
		}
	}
	panic("no start point")
}

func (f *Field) IsWall(x, y int) bool {
	return f.field[y*FIELD_X_MAX+x] != FIELD_NONE &&
		f.field[y*FIELD_X_MAX+x] != FIELD_HIDEPATH &&
		f.field[y*FIELD_X_MAX+x] != FIELD_BAR &&
		!f.IsItem(x, y)
}
func (f *Field) IsRidable(x, y int) bool {
	return f.field[y*FIELD_X_MAX+x] != FIELD_NONE &&
		f.field[y*FIELD_X_MAX+x] != FIELD_HIDEPATH &&
		!f.IsItem(x, y)
}

func (f *Field) IsSpike(x, y int) bool {
	return f.field[y*FIELD_X_MAX+x] == FIELD_SPIKE
}

func (f *Field) GetField(x, y int) FieldType {
	return f.field[y*FIELD_X_MAX+x]
}

func (f *Field) IsItem(x, y int) bool {
	return f.field[y*FIELD_X_MAX+x] >= FIELD_ITEM_BORDER &&
		f.field[y*FIELD_X_MAX+x] != FIELD_ITEM_STARTPOINT
}

func (f *Field) IsItemGettable(x, y int, playerData *PlayerData) bool {
	if !f.IsItem(x, y) {
		return false
	}
	if f.field[y*FIELD_X_MAX+x] == FIELD_ITEM_OMEGA && f.isHiddenSecret(playerData) {
		return false
	}
	return true
}

func (f *Field) isHiddenSecret(playerData *PlayerData) bool {
	return playerData.GetItemCount() < 15
}

func (f *Field) EraseField(x, y int) {
	f.field[y*FIELD_X_MAX+x] = FIELD_NONE
}

func (f *Field) Draw(game *Game, viewPosition Position) {
	parts := []imgPart{}
	vx, vy := viewPosition.X, viewPosition.Y
	ofs_x := CHAR_SIZE - vx%CHAR_SIZE
	ofs_y := CHAR_SIZE - vy%CHAR_SIZE
	for xx := -(g_width/CHAR_SIZE/2 + 2); xx < (g_width/CHAR_SIZE/2 + 2); xx++ {
		fx := xx + vx/CHAR_SIZE
		if fx < 0 || fx >= FIELD_X_MAX {
			continue
		}
		for yy := -(g_height/CHAR_SIZE/2 + 2); yy < (g_height/CHAR_SIZE/2 + 2); yy++ {
			fy := yy + vy/CHAR_SIZE
			if fy < 0 || fy >= FIELD_Y_MAX {
				continue
			}

			gy := (f.timer / 10) % 4
			gx := int(f.field[fy*FIELD_X_MAX+fx])

			if f.IsItem(fx, fy) {
				gx = gx - (int(FIELD_ITEM_BORDER) + 1)
				gy = 4 + gx/16
				gx = gx % 16
			}

			if f.isHiddenSecret(game.playerData) && f.field[fy*FIELD_X_MAX+fx] == FIELD_ITEM_OMEGA {
				continue
			}

			parts = append(parts, imgPart{
				(xx+12)*CHAR_SIZE + ofs_x + GRAPHIC_OFFSET_X + (g_width-320)/2,
				(yy+8)*CHAR_SIZE + ofs_y + GRAPHIC_OFFSET_Y + (g_height-240)/2,
				gx * 16, gy * 16, 16, 16,
			})

		}
	}
	game.DrawParts("ino", parts)
}
