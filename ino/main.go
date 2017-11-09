package ino

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/go-inovation/ino/internal/input"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	Title        = "Inovation 2007 (Go version)"
)

const (
	ENDINGMAIN_STATE_STAFFROLL = iota
	ENDINGMAIN_STATE_RESULT
)

type GameStateMsg int

const (
	GAMESTATE_MSG_NONE GameStateMsg = iota
	GAMESTATE_MSG_REQ_TITLE
	GAMESTATE_MSG_REQ_GAME
	GAMESTATE_MSG_REQ_OPENING
	GAMESTATE_MSG_REQ_ENDING
	GAMESTATE_MSG_REQ_SECRET1
	GAMESTATE_MSG_REQ_SECRET2
)

type TitleMain struct {
	gameStateMsg  GameStateMsg
	timer         int
	offsetX       int
	offsetY       int
	lunkerMode    bool
	lunkerCommand int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (t *TitleMain) Update(game *Game) {
	t.timer++
	if t.timer%5 == 0 {
		t.offsetX = rand.Intn(5) - 3
		t.offsetY = rand.Intn(5) - 3
	}

	if (input.Current().IsActionKeyJustPressed() || input.Current().IsSpaceJustTouched()) && t.timer > 5 {
		t.gameStateMsg = GAMESTATE_MSG_REQ_OPENING

		if t.lunkerMode {
			game.gameData = NewGameData(GAMEMODE_LUNKER)
		} else {
			game.gameData = NewGameData(GAMEMODE_NORMAL)
		}
	}

	// ランカー・モード・コマンド
	switch t.lunkerCommand {
	case 0, 1, 2, 6:
		if input.Current().IsKeyJustPressed(ebiten.KeyLeft) {
			t.lunkerCommand++
		} else if input.Current().IsKeyJustPressed(ebiten.KeyRight) || input.Current().IsKeyJustPressed(ebiten.KeyUp) || input.Current().IsKeyJustPressed(ebiten.KeyDown) {
			t.lunkerCommand = 0
		}
	case 3, 4, 5, 7:
		if input.Current().IsKeyJustPressed(ebiten.KeyRight) {
			t.lunkerCommand++
		} else if input.Current().IsKeyJustPressed(ebiten.KeyLeft) || input.Current().IsKeyJustPressed(ebiten.KeyUp) || input.Current().IsKeyJustPressed(ebiten.KeyDown) {
			t.lunkerCommand = 0
		}
	default:
		break
	}
	if t.lunkerCommand > 7 {
		t.lunkerCommand = 0
		t.lunkerMode = !t.lunkerMode
	}
}

func (t *TitleMain) Draw(game *Game) {
	if t.lunkerMode {
		game.Draw("bg", 0, 0, 0, 240, ScreenWidth, ScreenHeight)
		game.Draw("msg", (ScreenWidth-256)/2+t.offsetX, 160+t.offsetY+(ScreenHeight-240)/2, 0, 64, 256, 16)
	} else {
		game.Draw("bg", 0, 0, 0, 0, ScreenWidth, ScreenHeight)
		sy := 64 + 16
		if input.Current().IsTouchEnabled() {
			sy = 64 - 16
		}
		game.Draw("msg", (ScreenWidth-256)/2+t.offsetX, 160+t.offsetY+(ScreenHeight-240)/2, 0, sy, 256, 16)
	}
	game.Draw("msg", (ScreenWidth-256)/2, 32+(ScreenHeight-240)/2, 0, 0, 256, 48)
}

func (t *TitleMain) Msg() GameStateMsg {
	return t.gameStateMsg
}

type OpeningMain struct {
	gameStateMsg GameStateMsg
	timer        int
}

const (
	OPENING_SCROLL_LEN   = 416
	OPENING_SCROLL_SPEED = 3
)

func (o *OpeningMain) Update(game *Game) {
	o.timer++

	if input.Current().IsActionKeyPressed() || input.Current().IsSpaceTouched() {
		o.timer += 20
	}
	if o.timer/OPENING_SCROLL_SPEED > OPENING_SCROLL_LEN+ScreenHeight {
		o.gameStateMsg = GAMESTATE_MSG_REQ_GAME
		PauseBGM()
	}
}

func (o *OpeningMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 480, 320, 240)
	game.Draw("msg", (ScreenWidth-256)/2, ScreenHeight-(o.timer/OPENING_SCROLL_SPEED), 0, 160, 256, OPENING_SCROLL_LEN)
}

func (o *OpeningMain) Msg() GameStateMsg {
	return o.gameStateMsg
}

type EndingMain struct {
	gameStateMsg   GameStateMsg
	timer          int
	bgmFadingTimer int
	state          int
}

const (
	ENDING_SCROLL_LEN   = 1088
	ENDING_SCROLL_SPEED = 3
)

func (e *EndingMain) Update(game *Game) {
	e.timer++
	switch e.state {
	case ENDINGMAIN_STATE_STAFFROLL:
		if input.Current().IsActionKeyPressed() || input.Current().IsSpaceTouched() {
			e.timer += 20
		}
		if e.timer/ENDING_SCROLL_SPEED > ENDING_SCROLL_LEN+ScreenHeight {
			e.timer = 0
			e.state = ENDINGMAIN_STATE_RESULT
		}
	case ENDINGMAIN_STATE_RESULT:
		const max = 5 * ebiten.FPS
		e.bgmFadingTimer++
		switch {
		case e.bgmFadingTimer == max:
			PauseBGM()
		case e.bgmFadingTimer < max:
			vol := 1 - (float64(e.bgmFadingTimer) / max)
			SetBGMVolume(vol)
		}
		if (input.Current().IsActionKeyJustPressed() || input.Current().IsSpaceJustTouched()) && e.timer > 5 {
			// 条件を満たしていると隠し画面へ
			if game.gameData.IsGetOmega() {
				if game.gameData.lunkerMode {
					e.gameStateMsg = GAMESTATE_MSG_REQ_SECRET2
				} else {
					e.gameStateMsg = GAMESTATE_MSG_REQ_SECRET1
				}
				return
			}
			e.gameStateMsg = GAMESTATE_MSG_REQ_TITLE
			PauseBGM()
		}
	}
}

func (e *EndingMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 480, 320, 240)

	switch e.state {
	case ENDINGMAIN_STATE_STAFFROLL:
		game.Draw("msg", (ScreenWidth-256)/2, ScreenHeight-(e.timer/ENDING_SCROLL_SPEED), 0, 576, 256, ENDING_SCROLL_LEN)
	case ENDINGMAIN_STATE_RESULT:
		game.Draw("msg", (ScreenWidth-256)/2, (ScreenHeight-160)/2, 0, 1664, 256, 160)
		game.DrawNumber(game.gameData.GetItemCount(), (ScreenWidth-10*0)/2, (ScreenHeight-160)/2+13*5+2)
		game.DrawNumber(game.gameData.TimeInSecond(), (ScreenWidth-13)/2, (ScreenHeight-160)/2+13*8+2)
	}
}

func (e *EndingMain) Msg() GameStateMsg {
	return e.gameStateMsg
}

type SecretMain struct {
	gameStateMsg GameStateMsg
	timer        int
	number       int
}

func NewSecretMain(number int) *SecretMain {
	return &SecretMain{
		number: number,
	}
}

func (s *SecretMain) Update(game *Game) {
	s.timer++
	if (input.Current().IsActionKeyJustPressed() || input.Current().IsSpaceJustTouched()) && s.timer > 5 {
		s.gameStateMsg = GAMESTATE_MSG_REQ_TITLE
	}
}

func (s *SecretMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 240, 320, 240)
	if s.number == 1 {
		game.Draw("msg", (ScreenWidth-256)/2, (ScreenHeight-96)/2, 0, 2048-96*2, 256, 96)
		return
	}
	game.Draw("msg", (ScreenWidth-256)/2, (ScreenHeight-96)/2, 0, 2048-96, 256, 96)
}

func (s *SecretMain) Msg() GameStateMsg {
	return s.gameStateMsg
}

type GameMain struct {
	gameStateMsg GameStateMsg
	player       *Player
}

func NewGameMain(game *Game) *GameMain {
	g := &GameMain{
		player: NewPlayer(game.gameData),
	}
	return g
}

func (g *GameMain) Update(game *Game) {
	g.gameStateMsg = g.player.Update()
}

func (g *GameMain) Draw(game *Game) {
	if game.gameData.lunkerMode {
		game.Draw("bg", 0, 0, 0, 240, ScreenWidth, ScreenHeight)
	} else {
		game.Draw("bg", 0, 0, 0, 0, ScreenWidth, ScreenHeight)
	}
	g.player.Draw(game)
	if input.Current().IsTouchEnabled() {
		game.DrawTouchButtons()
	}
}

func (g *GameMain) Msg() GameStateMsg {
	return g.gameStateMsg
}

type GameState interface {
	Update(g *Game) // TODO: Should return errors
	Draw(g *Game)
	Msg() GameStateMsg
}
