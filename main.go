package inovation5

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

const g_width = 320
const g_height = 240
const CHAR_SIZE = 16

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
			game.playerData = NewPlayerData(GAMEMODE_LUNKER)
		} else {
			game.playerData = NewPlayerData(GAMEMODE_NORMAL)
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

func (t *TitleMain) GetMsg() GameStateMsg {
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
		// TODO(hajimehoshi): Stop BGM
	}
}

func (o *OpeningMain) Draw(game *Game) {
	game.Draw("bg", 0, 0, 0, 480, 320, 240)
	game.Draw("msg", (g_width-256)/2, g_height-(o.timer/OPENING_SCROLL_SPEED), 0, 160, 256, OPENING_SCROLL_LEN)
}

func (o *OpeningMain) GetMsg() GameStateMsg {
	return o.gameStateMsg
}

type EndingMain struct {
	gameStateMsg GameStateMsg
	timer        int
	state        int
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
			// TODO(hajimehoshi): Stop BGM with fade 5000
		}
	case ENDINGMAIN_STATE_RESULT:
		if input.IsActionKeyPushed() && e.timer > 5 {
			// 条件を満たしていると隠し画面へ
			if game.playerData.IsGetOmega() {
				if game.playerData.lunkerMode {
					e.gameStateMsg = GAMESTATE_MSG_REQ_SECRET2
				} else {
					e.gameStateMsg = GAMESTATE_MSG_REQ_SECRET1
				}
			} else {
				e.gameStateMsg = GAMESTATE_MSG_REQ_TITLE
			}
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

		game.DrawFont(strconv.Itoa(game.playerData.GetItemCount()), (g_width-10*0)/2, (g_height-160)/2+13*5+2)
		game.DrawFont(strconv.Itoa(game.playerData.playtime), (g_width-13)/2, (g_height-160)/2+13*8+2)
	}
}

func (e *EndingMain) GetMsg() GameStateMsg {
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

func (s *SecretMain) GetMsg() GameStateMsg {
	return s.gameStateMsg
}

type GameMain struct {
	gameStateMsg GameStateMsg
	player       *Player
	field        *Field
}

func NewGameMain(game *Game) *GameMain {
	f := NewField(field_data)
	g := &GameMain{
		player: NewPlayer(game.playerData, f),
		field:  f,
	}
	return g
}

func (g *GameMain) Update(game *Game) {
	g.field.Move()
	g.player.Move(g)
}

func (g *GameMain) Draw(game *Game) {
	if game.playerData.lunkerMode {
		game.Draw("bg", 0, 0, 0, 240, g_width, g_height)
	} else {
		game.Draw("bg", 0, 0, 0, 0, g_width, g_height)
	}

	p := g.player.view.GetPosition()
	g.field.Draw(game, Position{X: int(p.X), Y: int(p.Y)})
	g.player.Draw(game)
}

func (g *GameMain) GetMsg() GameStateMsg {
	return g.gameStateMsg
}

func (g *GameMain) SetMsg(msg GameStateMsg) {
	g.gameStateMsg = msg
}

type GameState interface {
	Update(g *Game)
	Draw(g *Game)
	GetMsg() GameStateMsg
}

type Game struct {
	gameState  GameState
	playerData *PlayerData
	img        map[string]*ebiten.Image
	font       *Font
	screen     *ebiten.Image
}

func NewGame() *Game {
	return &Game{
		playerData: NewPlayerData(GAMEMODE_NORMAL),
	}
}

func (g *Game) Start() error {
	return ebiten.Run(g.Loop, g_width, g_height, 2, "Inovation 5")
}

func (g *Game) Loop(screen *ebiten.Image) error {
	input.Update()
	g.screen = screen

	if g.gameState == nil {
		g.gameState = &TitleMain{}
	} else {
		switch g.gameState.GetMsg() {
		case GAMESTATE_MSG_REQ_TITLE:
			g.gameState = &TitleMain{}
			break
		case GAMESTATE_MSG_REQ_OPENING:
			// TODO(hajimehoshi): Play BGM 'bgm1'
			g.gameState = &OpeningMain{}
			break
		case GAMESTATE_MSG_REQ_GAME:
			g.gameState = NewGameMain(g)
			break
		case GAMESTATE_MSG_REQ_ENDING:
			// TODO(hajimehoshi): Play BGM 'bgm1'
			g.gameState = &EndingMain{}
			break
		case GAMESTATE_MSG_REQ_SECRET1:
			// TODO(hajimehoshi): Play BGM 'bgm1'
			g.gameState = NewSecretMain(1)
			break
		case GAMESTATE_MSG_REQ_SECRET2:
			// TODO(hajimehoshi): Play BGM 'bgm1'
			g.gameState = NewSecretMain(2)
			break
		}
	}
	g.gameState.Update(g)
	g.gameState.Draw(g)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n%.2f", ebiten.CurrentFPS()))
	return nil
}

func (g *Game) FillRect(x, y, w, h int, clr color.Color) error {
	return g.screen.DrawRect(x, y, w, h, clr)
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
	// TODO(hajimehoshi): Load SE
	// TODO(hajimehoshi): Load BGM

	const imgDir = "../resource/image/color"

	game := NewGame()
	game.img = map[string]*ebiten.Image{}
	for _, f := range []string{"ino", "msg", "bg"} {
		var err error
		game.img[f], _, err = ebitenutil.NewImageFromFile(filepath.Join(imgDir, f+".png"), ebiten.FilterNearest)
		if err != nil {
			return err
		}
	}

	game.font = NewFont()
	if err := game.font.Load("../resource/font"); err != nil {
		return err
	}
	return game.Start()
}
