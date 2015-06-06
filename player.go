package inovation5

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type PlayerState int

const (
	PLAYERSTATE_START PlayerState = iota
	PLAYERSTATE_NORMAL
	PLAYERSTATE_ITEMGET
	PLAYERSTATE_MUTEKI
	PLAYERSTATE_DEAD
)

const (
	PLAYER_SPEED          = 2.0
	PLAYER_GRD_ACCRATIO   = 0.04
	PLAYER_AIR_ACCRATIO   = 0.01
	PLAYER_JUMP           = -4.0
	PLAYER_GRAVITY        = 0.2
	PLAYER_FALL_SPEEDMAX  = 4.0
	VIEW_DIRECTION_OFFSET = 30.0
	WAIT_TIMER_INTERVAL   = 10
	LIFE_RATIO            = 400
	MUTEKI_INTERVAL       = 50
	START_WAIT_INTERVAL   = 50

	LUNKER_JUMP_DAMAGE1 = 40.0
	LUNKER_JUMP_DAMAGE2 = 96.0
)

type View struct {
	position PositionF
}

func (v *View) ToScreenPosition(p Position) Position {
	x := p.X - int(v.position.X) + g_width/2
	y := p.Y - int(v.position.Y) + g_height/2
	return Position{x, y}
}

func (v *View) GetPosition() PositionF {
	return v.position
}

func (v *View) SetPosition(p PositionF) {
	v.position = p
}

type Player struct {
	life        int
	jumpCnt     int
	timer       int
	position    PositionF // TODO(hajimehoshi): This can be a Position.
	speed       PositionF
	direction   int
	jumpedPoint PositionF
	state       PlayerState
	itemGet     int
	waitTimer   int
	game        *Game
	view        *View
}

func NewPlayer(g *Game) *Player {
	return &Player{game: g}
}

func (p *Player) Initialize() {
	*p = Player{game: p.game}
	p.state = PLAYERSTATE_NORMAL
	p.life = p.game.playerData.lifeMax * LIFE_RATIO

	startPoint := p.game.field.GetStartPoint()
	startPointF := PositionF{float64(startPoint.X), float64(startPoint.Y)}
	p.position = startPointF
	p.jumpedPoint = startPointF

	p.view = &View{}
	p.view.SetPosition(p.position)
	// TODO(hajimehoshi): Play BGM 'bgm0'
}

func (p *Player) OnWall() bool {
	if p.toFieldOfsY() > CHAR_SIZE/4 {
		return false
	}
	if p.game.field.IsRidable(p.toFieldX(), p.toFieldY()+1) && p.toFieldOfsX() < CHAR_SIZE*7/8 {
		return true
	}
	if p.game.field.IsRidable(p.toFieldX()+1, p.toFieldY()+1) && p.toFieldOfsX() > CHAR_SIZE/8 {
		return true
	}
	return false
}

func (p *Player) IsFallable() bool {
	if !p.OnWall() {
		return false
	}
	if p.game.field.IsWall(p.toFieldX(), p.toFieldY()+1) && p.toFieldOfsX() < CHAR_SIZE*7/8 {
		return false
	}
	if p.game.field.IsWall(p.toFieldX()+1, p.toFieldY()+1) && p.toFieldOfsX() > CHAR_SIZE/8 {
		return false
	}
	return true
}

func (p *Player) IsUpperWallBoth() bool {
	if p.toFieldOfsY() < CHAR_SIZE/2 {
		return false
	}
	if p.game.field.IsWall(p.toFieldX(), p.toFieldY()) && p.game.field.IsWall(p.toFieldX()+1, p.toFieldY()) {
		return true
	}
	return false
}

func (p *Player) IsUpperWall() bool {
	if p.toFieldOfsY() < CHAR_SIZE/2 {
		return false
	}
	if p.game.field.IsWall(p.toFieldX(), p.toFieldY()) && p.toFieldOfsX() < CHAR_SIZE*7/8 {
		return true
	}
	if p.game.field.IsWall(p.toFieldX()+1, p.toFieldY()) && p.toFieldOfsX() > CHAR_SIZE/8 {
		return true
	}
	return false
}
func (p *Player) IsLeftWall() bool {
	if p.game.field.IsWall(p.toFieldX(), p.toFieldY()) {
		return true
	}
	if p.game.field.IsWall(p.toFieldX(), p.toFieldY()+1) && p.toFieldOfsY() > CHAR_SIZE/8 {
		return true
	}
	return false
}
func (p *Player) IsRightWall() bool {
	if p.game.field.IsWall(p.toFieldX()+1, p.toFieldY()) {
		return true
	}
	if p.game.field.IsWall(p.toFieldX()+1, p.toFieldY()+1) && p.toFieldOfsY() > CHAR_SIZE/8 {
		return true
	}
	return false
}

func (p *Player) normalizeToRight() {
	p.position.X = float64(p.toFieldX() * CHAR_SIZE)
	p.speed.X = 0
}
func (p *Player) normalizeToLeft() {
	p.position.X = float64((p.toFieldX() + 1) * CHAR_SIZE)
	p.speed.X = 0
}
func (p *Player) normalizeToUpper() {
	if p.speed.Y < 0 {
		p.speed.Y = 0
	}
	p.position.Y = float64(CHAR_SIZE * (p.toFieldY() + 1))
}

func (p *Player) toFieldX() int {
	return int(p.position.X) / CHAR_SIZE
}

func (p *Player) toFieldY() int {
	return int(p.position.Y) / CHAR_SIZE
}

func (p *Player) toFieldOfsX() int {
	return int(p.position.X) % CHAR_SIZE
}
func (p *Player) toFieldOfsY() int {
	return int(p.position.Y) % CHAR_SIZE
}

func (p *Player) Move() {
	switch p.state {
	case PLAYERSTATE_START:
		p.waitTimer++
		if p.waitTimer > START_WAIT_INTERVAL {
			p.state = PLAYERSTATE_NORMAL
		}

	case PLAYERSTATE_NORMAL:
		p.moveByInput()
		p.moveNormal()
		if p.life < p.game.playerData.lifeMax*LIFE_RATIO {
			var o_life = p.life
			p.life++
			if (p.life / LIFE_RATIO) != (o_life / LIFE_RATIO) {
				// TODO(hajimehoshi): Play SE 'heal'
			}
		}

	case PLAYERSTATE_ITEMGET:
		p.MoveItemGet()
		if p.state != PLAYERSTATE_ITEMGET {
			if p.game.playerData.IsGameClear() {
				p.game.gameState.SetMsg(GAMESTATE_MSG_REQ_ENDING)
			}
		}

	case PLAYERSTATE_MUTEKI:
		p.moveByInput()
		p.moveNormal()
		p.waitTimer++
		if p.waitTimer > MUTEKI_INTERVAL {
			p.state = PLAYERSTATE_NORMAL
		}

	case PLAYERSTATE_DEAD:
		p.moveNormal()
		// TODO(hajimehoshi): Stop BGM
		if p.game.key.IsActionKeyPressed() && p.waitTimer > 15 {
			p.game.gameState.SetMsg(GAMESTATE_MSG_REQ_TITLE)
		}
	}
	if p.life < LIFE_RATIO {
		if p.state != PLAYERSTATE_DEAD {
			p.waitTimer = 0
		}
		p.state = PLAYERSTATE_DEAD
		p.direction = 0
		p.waitTimer++
	}
}

func (p *Player) moveNormal() {
	p.timer++
	p.game.playerData.playtime = (p.timer / 50)

	// 移動＆落下
	p.speed.Y += PLAYER_GRAVITY
	p.position.X += p.speed.X
	p.position.Y += p.speed.Y

	if p.speed.Y > PLAYER_FALL_SPEEDMAX {
		p.speed.Y = PLAYER_FALL_SPEEDMAX
	}

	if p.state == PLAYERSTATE_NORMAL {
		p.CheckCollision()
	}

	// ATARI判定
	hitLeft := false
	hitRight := false
	hitUpper := false
	if p.OnWall() && p.speed.Y >= 0 {
		if p.game.playerData.lunkerMode {
			if p.position.Y-p.jumpedPoint.Y > LUNKER_JUMP_DAMAGE1 {
				p.state = PLAYERSTATE_MUTEKI
				p.waitTimer = 0
				p.life -= LIFE_RATIO
				// TODO(hajimehoshi): Play SE 'damage'
			}
			if p.position.Y-p.jumpedPoint.Y > LUNKER_JUMP_DAMAGE2 {
				p.state = PLAYERSTATE_MUTEKI
				p.waitTimer = 0
				p.life -= LIFE_RATIO * 99
				// TODO(hajimehoshi): Play SE 'damage'
			}
		}

		if p.game.key.IsActionKeyPressed() && p.game.key.IsKeyPressed(ebiten.KeyDown) && p.IsFallable() {
			// 落下
		} else {
			if p.speed.Y > 0 {
				p.speed.Y = 0
			}
			p.position.Y = float64(CHAR_SIZE * p.toFieldY())
			p.jumpCnt = 0
		}

		p.jumpedPoint.X = p.position.X
		p.jumpedPoint.Y = p.position.Y
	}
	if p.IsLeftWall() && p.speed.X < 0 {
		hitLeft = true
	}
	if p.IsRightWall() && p.speed.X > 0 {
		hitRight = true
	}
	if p.IsUpperWall() && p.speed.Y <= 0 {
		hitUpper = true
	}

	if hitUpper && !hitLeft && !hitRight {
		p.normalizeToUpper()
	}
	if !hitUpper && hitLeft {
		p.normalizeToLeft()
	}
	if !hitUpper && hitRight {
		p.normalizeToRight()
	}
	if hitUpper && hitRight {
		if p.IsUpperWallBoth() {
			p.normalizeToUpper()
		} else {
			if p.toFieldOfsX() > p.toFieldOfsY() {
				p.normalizeToRight()
			} else {
				p.normalizeToUpper()
			}
		}
	}
	if hitUpper && hitLeft {
		if p.IsUpperWallBoth() {
			p.normalizeToUpper()
		} else {
			if CHAR_SIZE-p.toFieldOfsX() > p.toFieldOfsY() {
				p.normalizeToLeft()
			} else {
				p.normalizeToUpper()
			}
		}
	}

	// 床特殊効果
	switch p.GetOnField() {
	case FIELD_SCROLL_L:
		p.speed.X = p.speed.X*(1.0-PLAYER_GRD_ACCRATIO) + float64(p.direction*PLAYER_SPEED-SCROLLPANEL_SPEED)*PLAYER_GRD_ACCRATIO
	case FIELD_SCROLL_R:
		p.speed.X = p.speed.X*(1.0-PLAYER_GRD_ACCRATIO) + float64(p.direction*PLAYER_SPEED+SCROLLPANEL_SPEED)*PLAYER_GRD_ACCRATIO
	case FIELD_SLIP:
		// Do nothing
	case FIELD_NONE:
		p.speed.X = p.speed.X*(1.0-PLAYER_AIR_ACCRATIO) + float64(p.direction*PLAYER_SPEED)*PLAYER_AIR_ACCRATIO
	default:
		p.speed.X = p.speed.X*(1.0-PLAYER_GRD_ACCRATIO) + float64(p.direction*PLAYER_SPEED)*PLAYER_GRD_ACCRATIO
	}

	// ビューの更新
	v := p.view.GetPosition()
	v.X = v.X*0.95 + (p.position.X+p.speed.X*VIEW_DIRECTION_OFFSET)*0.05
	v.Y = v.Y*0.95 + p.position.Y*0.05
	p.view.SetPosition(v)
}

func (p *Player) MoveItemGet() {
	if p.waitTimer < WAIT_TIMER_INTERVAL {
		p.waitTimer++
		return
	}
	if p.game.key.IsActionKeyPushed() {
		p.state = PLAYERSTATE_NORMAL
		// TODO(hajimehoshi): Play BGM 'bgm0'
	}
}

func (p *Player) moveByInput() {
	if p.game.key.IsKeyPressed(ebiten.KeyLeft) {
		p.direction = -1
	}
	if p.game.key.IsKeyPressed(ebiten.KeyRight) {
		p.direction = 1
	}

	if p.game.key.IsActionKeyPushed() {
		if ((p.game.playerData.jumpMax > p.jumpCnt) || p.OnWall()) && !p.game.key.IsKeyPressed(ebiten.KeyDown) {
			p.speed.Y = PLAYER_JUMP // ジャンプ
			if !p.OnWall() {
				p.jumpCnt++
			}

			if math.Abs(p.speed.X) < 0.1 {
				if p.speed.X < 0 {
					p.speed.X -= 0.02
				}
				if p.speed.X > 0 {
					p.speed.X += 0.02
				}
			}

			// TODO(hajimehoshi): Play SE 'jump'

			p.jumpedPoint = p.position
		}
	}
}

func (p *Player) CheckCollision() {
	for xx := 0; xx < 2; xx++ {
		for yy := 0; yy < 2; yy++ {
			// アイテム獲得(STATE_ITEMGETへ遷移)
			if p.game.field.IsItem(p.toFieldX()+xx, p.toFieldY()+yy) {
				// 隠しアイテムは条件が必要
				if !p.game.field.IsItemGettable(p.toFieldX()+xx, p.toFieldY()+yy, p.game.playerData) {
					continue
				}

				p.state = PLAYERSTATE_ITEMGET

				// アイテム効果
				p.itemGet = p.game.field.GetField(p.toFieldX()+xx, p.toFieldY()+yy)
				switch p.game.field.GetField(p.toFieldX()+xx, p.toFieldY()+yy) {
				case FIELD_ITEM_POWERUP:
					p.game.playerData.jumpMax++
				case FIELD_ITEM_LIFE:
					p.game.playerData.lifeMax++
					p.life = p.game.playerData.lifeMax * LIFE_RATIO
				default:
					p.game.playerData.itemGetFlags[p.itemGet] = true
				}
				p.game.field.EraseField(p.toFieldX()+xx, p.toFieldY()+yy)
				p.waitTimer = 0

				// TODO(hajimehoshi): Stop BGM
				if p.game.playerData.IsItemForClear(p.itemGet) || p.itemGet == FIELD_ITEM_POWERUP {
					// TODO(hajimehoshi): Play SE 'itemget'
				} else {
					// TODO(hajimehoshi): Play SE 'itemget2'
				}
				return
			}
			// トゲ(ダメージ)
			if p.game.field.IsSpike(p.toFieldX()+xx, p.toFieldY()+yy) {
				p.state = PLAYERSTATE_MUTEKI
				p.waitTimer = 0
				p.life -= LIFE_RATIO
				p.speed.Y = PLAYER_JUMP
				p.jumpCnt = -1 // ダメージ・エキストラジャンプ

				// TODO(hajimehoshi): Play SE 'damage'

				return
			}
		}
	}
}

func (p *Player) GetOnField() int {
	if !p.OnWall() {
		return FIELD_NONE
	}
	if p.toFieldOfsX() < CHAR_SIZE/2 {
		if p.game.field.IsRidable(p.toFieldX(), p.toFieldY()+1) {
			return p.game.field.GetField(p.toFieldX(), p.toFieldY()+1)
		} else {
			return p.game.field.GetField(p.toFieldX()+1, p.toFieldY()+1)
		}
	} else {
		if p.game.field.IsRidable(p.toFieldX()+1, p.toFieldY()+1) {
			return p.game.field.GetField(p.toFieldX()+1, p.toFieldY()+1)
		} else {
			return p.game.field.GetField(p.toFieldX(), p.toFieldY()+1)
		}
	}
}

func (p *Player) Draw() {
	v := p.view.ToScreenPosition(Position{int(p.position.X), int(p.position.Y)})
	if p.state == PLAYERSTATE_DEAD { // 死亡
		anime := (p.timer / 6) % 4
		if p.game.playerData.lunkerMode {
			p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*(2+anime), 128+CHAR_SIZE*2, CHAR_SIZE, CHAR_SIZE)
		} else {
			p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*(2+anime), 128, CHAR_SIZE, CHAR_SIZE)
		}
	} else { // 生存
		if p.state != PLAYERSTATE_MUTEKI || p.timer%10 < 5 {
			anime := (p.timer / 6) % 2
			if !p.OnWall() {
				anime = 0
			}
			if p.direction < 0 {
				if p.game.playerData.lunkerMode {
					p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*anime, 128+CHAR_SIZE*2, CHAR_SIZE, CHAR_SIZE)
				} else {
					p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*anime, 128, CHAR_SIZE, CHAR_SIZE)
				}
			} else {
				if p.game.playerData.lunkerMode {
					p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*anime, 128+CHAR_SIZE*3, CHAR_SIZE, CHAR_SIZE)
				} else {
					p.game.Draw("ino", v.X, v.Y, CHAR_SIZE*anime, 128+CHAR_SIZE, CHAR_SIZE, CHAR_SIZE)
				}
			}
		}
	}

	// ライフ表示
	for t := 0; t < p.game.playerData.lifeMax; t++ {
		if p.life < LIFE_RATIO*2 && p.timer%10 < 5 && p.game.playerData.lifeMax > 1 {
			continue
		}

		if p.life >= (t+1)*LIFE_RATIO {
			p.game.Draw("ino", CHAR_SIZE*t, 0, CHAR_SIZE*3, 128+CHAR_SIZE*1, CHAR_SIZE, CHAR_SIZE)
		} else {
			p.game.Draw("ino", CHAR_SIZE*t, 0, CHAR_SIZE*4, 128+CHAR_SIZE*1, CHAR_SIZE, CHAR_SIZE)
		}
	}

	// 取ったアイテム一覧
	for t := FIELD_ITEM_FUJI; t < FIELD_ITEM_MAX; t++ {
		if !p.game.playerData.itemGetFlags[t] {
			p.game.Draw("ino", g_width-CHAR_SIZE/4*(FIELD_ITEM_MAX-2-t), 0, // 無
				CHAR_SIZE*5, 128+CHAR_SIZE, CHAR_SIZE/4, CHAR_SIZE/2)
		} else {
			if p.game.playerData.IsItemForClear(t) {
				// クリア条件アイテムは専用グラフィック
				for i, c := range clearFlagItems {
					if c == t {
						p.game.Draw("ino", g_width-CHAR_SIZE/4*(FIELD_ITEM_MAX-2-t), 0,
							CHAR_SIZE*5+CHAR_SIZE/4*(i+2), 128+CHAR_SIZE, CHAR_SIZE/4, CHAR_SIZE/2)
					}
				}
			} else {
				p.game.Draw("ino", g_width-CHAR_SIZE/4*(FIELD_ITEM_MAX-2-t), 0, // 有
					CHAR_SIZE*5+CHAR_SIZE/4, 128+CHAR_SIZE, CHAR_SIZE/4, CHAR_SIZE/2)
			}
		}
	}

	// アイテム獲得メッセージ
	if p.state == PLAYERSTATE_ITEMGET {
		t := WAIT_TIMER_INTERVAL - p.waitTimer
		p.game.Draw("msg", (g_width-256)/2, (g_height-96)/2-t*t+24,
			256, 96*(p.itemGet-FIELD_ITEM_BORDER-1), 256, 96)
		p.game.FillRect((g_width-32)/2, (g_height-96)/2-t*t-24, 32, 32, color.RGBA{0, 0, 0, 255})
		p.game.FillRect((g_width-32)/2+2, (g_height-96)/2-t*t-24+2, 32-4, 32-4, color.RGBA{255, 255, 255, 255})

		it := p.itemGet - (FIELD_ITEM_BORDER + 1)
		p.game.Draw("ino", (g_width-16)/2, (g_height-96)/2-t*t-16,
			(it%16)*CHAR_SIZE, (it/16+4)*CHAR_SIZE, CHAR_SIZE, CHAR_SIZE)
	}

	// ゲーム開始メッセージ
	if p.state == PLAYERSTATE_START {
		p.game.Draw("msg", (g_width-256)/2, 64+(g_height-240)/2, 0, 96, 256, 32)
	}

	// ゲームオーバーメッセージ
	if p.state == PLAYERSTATE_DEAD {
		p.game.Draw("msg", (g_width-256)/2, 64+(g_height-240)/2, 0, 128, 256, 32)
	}
}
