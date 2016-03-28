package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	g_width   = 320
	g_height  = 240
	CHAR_SIZE = 16
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

	if input.IsActionKeyPushed() && t.timer > 5 {
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
		if input.IsKeyPushed(ebiten.KeyLeft) {
			t.lunkerCommand++
		} else if input.IsKeyPushed(ebiten.KeyRight) || input.IsKeyPushed(ebiten.KeyUp) || input.IsKeyPushed(ebiten.KeyDown) {
			t.lunkerCommand = 0
		}
	case 3, 4, 5, 7:
		if input.IsKeyPushed(ebiten.KeyRight) {
			t.lunkerCommand++
		} else if input.IsKeyPushed(ebiten.KeyLeft) || input.IsKeyPushed(ebiten.KeyUp) || input.IsKeyPushed(ebiten.KeyDown) {
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
		game.Draw("bg", 0, 0, 0, 240, g_width, g_height)
		game.Draw("msg", (g_width-256)/2+t.offsetX, 160+t.offsetY+(g_height-240)/2, 0, 64, 256, 16)
	} else {
		game.Draw("bg", 0, 0, 0, 0, g_width, g_height)
		game.Draw("msg", (g_width-256)/2+t.offsetX, 160+t.offsetY+(g_height-240)/2, 0, 64+16, 256, 16)
	}
	game.Draw("msg", (g_width-256)/2, 32+(g_height-240)/2, 0, 0, 256, 64)
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

	if input.IsActionKeyPressed() {
		o.timer += 20
	}
	if o.timer/OPENING_SCROLL_SPEED > OPENING_SCROLL_LEN+g_height {
		o.gameStateMsg = GAMESTATE_MSG_REQ_GAME
		StopBGM()
	}
}

func (o *OpeningMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 480, 320, 240)
	game.Draw("msg", (g_width-256)/2, g_height-(o.timer/OPENING_SCROLL_SPEED), 0, 160, 256, OPENING_SCROLL_LEN)
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
		if input.IsActionKeyPressed() {
			e.timer += 20
		}
		if e.timer/ENDING_SCROLL_SPEED > ENDING_SCROLL_LEN+g_height {
			e.timer = 0
			e.state = ENDINGMAIN_STATE_RESULT
		}
	case ENDINGMAIN_STATE_RESULT:
		const max = 5 * ebiten.FPS
		e.bgmFadingTimer++
		switch {
		case e.bgmFadingTimer == max:
			StopBGM()
		case e.bgmFadingTimer < max:
			vol := 1 - (float64(e.bgmFadingTimer) / max)
			SetBGMVolume(vol)
		}
		if input.IsActionKeyPushed() && e.timer > 5 {
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
			StopBGM()
		}
	}
}

func (e *EndingMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 480, 320, 240)

	switch e.state {
	case ENDINGMAIN_STATE_STAFFROLL:
		game.Draw("msg", (g_width-256)/2, g_height-(e.timer/ENDING_SCROLL_SPEED), 0, 576, 256, ENDING_SCROLL_LEN)
	case ENDINGMAIN_STATE_RESULT:
		game.Draw("msg", (g_width-256)/2, (g_height-160)/2, 0, 1664, 256, 160)

		game.DrawFont(strconv.Itoa(game.gameData.GetItemCount()), (g_width-10*0)/2, (g_height-160)/2+13*5+2)
		game.DrawFont(strconv.Itoa(game.gameData.TimeInSecond()), (g_width-13)/2, (g_height-160)/2+13*8+2)
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
	if input.IsActionKeyPushed() && s.timer > 5 {
		s.gameStateMsg = GAMESTATE_MSG_REQ_TITLE
	}
}

func (s *SecretMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 240, 320, 240)

	if s.number == 1 {
		game.Draw("msg", (g_width-256)/2, (g_height-96)/2, 0, 2048-96*2, 256, 96)
	} else {
		game.Draw("msg", (g_width-256)/2, (g_height-96)/2, 0, 2048-96, 256, 96)
	}
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
		game.Draw("bg", 0, 0, 0, 240, g_width, g_height)
	} else {
		game.Draw("bg", 0, 0, 0, 0, g_width, g_height)
	}
	g.player.Draw(game)
}

func (g *GameMain) Msg() GameStateMsg {
	return g.gameStateMsg
}

type GameState interface {
	Update(g *Game)
	Draw(g *Game)
	Msg() GameStateMsg
}

type Game struct {
	gameState GameState
	gameData  *GameData
	img       map[string]*ebiten.Image
	font      *Font
	screen    *ebiten.Image
}

func (g *Game) Start() error {
	return ebiten.Run(g.Loop, g_width, g_height, 2, "Inovation (Go version)")
}

func (g *Game) Loop(screen *ebiten.Image) error {
	audioContext.Update()
	input.Update()
	g.screen = screen

	if g.gameState == nil {
		g.gameState = &TitleMain{}
	} else {
		switch g.gameState.Msg() {
		case GAMESTATE_MSG_REQ_TITLE:
			g.gameState = &TitleMain{}
		case GAMESTATE_MSG_REQ_OPENING:
			if err := PlayBGM(BGM1); err != nil {
				return err
			}
			g.gameState = &OpeningMain{}
		case GAMESTATE_MSG_REQ_GAME:
			g.gameState = NewGameMain(g)
		case GAMESTATE_MSG_REQ_ENDING:
			if err := PlayBGM(BGM1); err != nil {
				return err
			}
			g.gameState = &EndingMain{}
		case GAMESTATE_MSG_REQ_SECRET1:
			if err := PlayBGM(BGM1); err != nil {
				return err
			}
			g.gameState = NewSecretMain(1)
		case GAMESTATE_MSG_REQ_SECRET2:
			if err := PlayBGM(BGM1); err != nil {
				return err
			}
			g.gameState = NewSecretMain(2)
		}
	}
	g.gameState.Update(g)
	if !ebiten.IsRunningSlowly() {
		g.gameState.Draw(g)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n%.2f", ebiten.CurrentFPS()))
	return nil
}

var (
	imageEmpty *ebiten.Image
)

func init() {
	var err error
	imageEmpty, err = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

func toNRGBA(clr color.Color) (fr, fg, fb, fa float64) {
	r, g, b, a := clr.RGBA()
	fr = float64(r) / float64(a)
	fg = float64(g) / float64(a)
	fb = float64(b) / float64(a)
	fa = float64(a) / 0xff
	return
}

func (g *Game) FillRect(x, y, w, h int, clr color.Color) error {
	op := &ebiten.DrawImageOptions{}
	ew, eh := imageEmpty.Size()
	op.GeoM.Scale(float64(w)/float64(ew), float64(h)/float64(eh))
	cr, cg, cb, ca := toNRGBA(clr)
	if ca == 0 {
		return nil
	}
	op.ColorM.Scale(cr, cg, cb, ca)
	return g.screen.DrawImage(imageEmpty, op)
}

type imgPart struct {
	px, py, sx, sy, sw, sh int
}

type imgParts []imgPart

func (i imgParts) Len() int {
	return len(i)
}

func (i imgParts) Dst(idx int) (int, int, int, int) {
	p := i[idx]
	return p.px, p.py, p.px + p.sw, p.py + p.sh
}

func (i imgParts) Src(idx int) (int, int, int, int) {
	p := i[idx]
	return p.sx, p.sy, p.sx + p.sw, p.sy + p.sh
}

func (g *Game) Draw(key string, px, py, sx, sy, sw, sh int) error {
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = imgParts([]imgPart{
		imgPart{px, py, sx, sy, sw, sh},
	})
	return g.screen.DrawImage(g.img[key], op)
}

func (g *Game) DrawParts(key string, parts []imgPart) error {
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = imgParts(parts)
	return g.screen.DrawImage(g.img[key], op)
}

func (g *Game) DrawFont(msg string, x, y int) {
	for _, c := range msg {
		if c == 32 {
			x += 9
			continue
		}
		if img, ok := g.font.fonts[c]; ok {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			g.screen.DrawImage(img, op)
			w, _ := img.Size()
			x += w
		} else {
			x += 9
		}
	}
}

func Run() error {
	if err := initAudio(); err != nil {
		return err
	}

	const imgDir = "resource/image/color"
	game := &Game{
		img: map[string]*ebiten.Image{},
	}
	for _, f := range []string{"ino", "msg", "bg"} {
		var err error
		game.img[f], _, err = ebitenutil.NewImageFromFile(filepath.Join(imgDir, f+".png"), ebiten.FilterNearest)
		if err != nil {
			return err
		}
	}

	game.font = NewFont()
	if err := game.font.Load("resource/font"); err != nil {
		return err
	}
	return game.Start()
}

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}
