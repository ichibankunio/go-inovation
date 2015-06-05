package inovation5

import (
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

const ()

const (
	ENDINGMAIN_STATE_STAFFROLL = iota
	ENDINGMAIN_STATE_RESULT
)

type TitleMain struct {
	gameState     GameState
	timer         int
	offsetX       int
	offsetY       int
	lunkerMode    bool
	lunkerCommand int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (t *TitleMain) Update() {
	t.timer++
	if t.timer%5 == 0 {
		t.offsetX = int(rand.Float64()*3000/11)%5 - 3
		t.offsetY = int(rand.Float64()*3000/11)%5 - 3
	}

	key := t.gameState.game.key
	if key.IsActionKeyPushed() && t.timer > 5 {
		t.gameState.SetMsg(GAMESTATE_MSG_REQ_OPENING)

		if t.lunkerMode {
			t.gameState.game.playerData = NewPlayerData(GAMEMODE_LUNKER)
		} else {
			t.gameState.game.playerData = NewPlayerData(GAMEMODE_NORMAL)
		}
	}

	// ランカー・モード・コマンド
	switch t.lunkerCommand {
	case 0, 1, 2, 6:
		if key.IsKeyPushed(ebiten.KeyLeft) {
			t.lunkerCommand++
		} else if key.IsKeyPushed(ebiten.KeyRight) || key.IsKeyPushed(ebiten.KeyUp) || key.IsKeyPushed(ebiten.KeyDown) {
			t.lunkerCommand = 0
		}
	case 3, 4, 5, 7:
		if key.IsKeyPushed(ebiten.KeyRight) {
			t.lunkerCommand++
		} else if key.IsKeyPushed(ebiten.KeyLeft) || key.IsKeyPushed(ebiten.KeyUp) || key.IsKeyPushed(ebiten.KeyDown) {
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

func (t *TitleMain) Draw() {
	if t.lunkerMode {
		t.gameState.game.Draw("bg", 0, 0, 0, 240, 320, 240)
		t.gameState.game.Draw("msg", (g_width-256)/2+t.offsetX, 160+t.offsetY+(g_height-240)/2, 0, 64, 256, 16)
	} else {
		t.gameState.game.Draw("bg", 0, 0, 0, 0, 320, 240)
		t.gameState.game.Draw("msg", (g_width-256)/2+t.offsetX, 160+t.offsetY+(g_height-240)/2, 0, 64+16, 256, 16)
	}

	t.gameState.game.Draw("msg", (g_width-256)/2, 32+(g_height-240)/2, 0, 0, 256, 64)
}

func (t *TitleMain) GetMsg() int {
	return t.gameState.GetMsg()
}

func (t *TitleMain) SetMsg(msg int) {
	t.gameState.SetMsg(msg)
}

type OpeningMain struct {
	gameState GameState
	timer     int
}

const (
	OPENING_SCROLL_LEN   = 416
	OPENING_SCROLL_SPEED = 3
)

func NewOpeningMain(game *Game) *OpeningMain {
	// TODO(hajimehoshi): Play BGM 'bgm1'
	return &OpeningMain{
		gameState: GameState{game: game},
	}
}

func (o *OpeningMain) Update() {
	o.timer++

	if o.gameState.game.key.IsActionKeyPressed() {
		o.timer += 20
	}
	if o.timer/OPENING_SCROLL_SPEED > OPENING_SCROLL_LEN+g_height {
		o.gameState.SetMsg(GAMESTATE_MSG_REQ_GAME)
		// TODO(hajimehoshi): Stop BGM
	}
}

func (o *OpeningMain) Draw() {
	o.gameState.game.Draw("bg", 0, 0, 0, 480, 320, 240)
	o.gameState.game.Draw("msg", (g_width-256)/2, g_height-(o.timer/OPENING_SCROLL_SPEED), 0, 160, 256, OPENING_SCROLL_LEN)
}

func (o *OpeningMain) GetMsg() int {
	return o.gameState.GetMsg()
}

func (o *OpeningMain) SetMsg(msg int) {
	o.gameState.SetMsg(msg)
}

type EndingMain struct {
	gameState GameState
	timer     int
	state     int
}

const (
	ENDING_SCROLL_LEN   = 1088
	ENDING_SCROLL_SPEED = 3
)

func NewEndingMain(game *Game) *EndingMain {
	// TODO(hajimehoshi): Play BGM 'bgm1'
	return &EndingMain{
		gameState: GameState{game: game},
	}
}

func (e *EndingMain) Update() {
	e.timer++
	switch e.state {
	case ENDINGMAIN_STATE_STAFFROLL:
		if e.gameState.game.key.IsActionKeyPressed() {
			e.timer += 20
		}
		if e.timer/ENDING_SCROLL_SPEED > ENDING_SCROLL_LEN+g_height {
			e.timer = 0
			e.state = ENDINGMAIN_STATE_RESULT
			// TODO(hajimehoshi): Stop BGM with fade 5000
		}
	case ENDINGMAIN_STATE_RESULT:
		if e.gameState.game.key.IsActionKeyPushed() && e.timer > 5 {
			// 条件を満たしていると隠し画面へ
			if e.gameState.game.playerData.IsGetOmega() {
				if e.gameState.game.playerData.lunkerMode {
					e.gameState.SetMsg(GAMESTATE_MSG_REQ_SECRET2)
				} else {
					e.gameState.SetMsg(GAMESTATE_MSG_REQ_SECRET1)
				}
			} else {
				e.gameState.SetMsg(GAMESTATE_MSG_REQ_TITLE)
			}
		}
	}
}

func (e *EndingMain) Draw() {
	e.gameState.game.Draw("bg", 0, 0, 0, 480, 320, 240)

	switch e.state {
	case ENDINGMAIN_STATE_STAFFROLL:
		e.gameState.game.Draw("msg", (g_width-256)/2, g_height-(e.timer/ENDING_SCROLL_SPEED), 0, 576, 256, ENDING_SCROLL_LEN)
	case ENDINGMAIN_STATE_RESULT:
		e.gameState.game.Draw("msg", (g_width-256)/2, (g_height-160)/2, 0, 1664, 256, 160)

		e.gameState.game.DrawFont(strconv.Itoa(e.gameState.game.playerData.GetItemCount()), (g_width-10*0)/2, (g_height-160)/2+13*5+2)
		e.gameState.game.DrawFont(strconv.Itoa(e.gameState.game.playerData.playtime), (g_width-13)/2, (g_height-160)/2+13*8+2)
	}
}

func (e *EndingMain) GetMsg() int {
	return e.gameState.GetMsg()
}

func (e *EndingMain) SetMsg(msg int) {
	e.gameState.SetMsg(msg)
}

type SecretMain struct {
	gameState GameState
	timer     int
	number    int
}

func NewSecretMain(game *Game, number int) *SecretMain {
	// TODO(hajimehoshi): Play BGM 'bgm1'
	return &SecretMain{
		gameState: GameState{game: game},
		number:    number,
	}
}

func (s *SecretMain) Update() {
	s.timer++
	if s.gameState.game.key.IsActionKeyPushed() && s.timer > 5 {
		s.gameState.SetMsg(GAMESTATE_MSG_REQ_TITLE)
	}
}

func (s *SecretMain) Draw() {
	s.gameState.game.Draw("bg", 0, 0, 0, 240, 320, 240)

	if s.number == 1 {
		s.gameState.game.Draw("msg", (g_width-256)/2, (g_height-96)/2, 0, 2048-96*2, 256, 96)
	} else {
		s.gameState.game.Draw("msg", (g_width-256)/2, (g_height-96)/2, 0, 2048-96, 256, 96)
	}
}

func (s *SecretMain) GetMsg() int {
	return s.gameState.GetMsg()
}

func (s *SecretMain) SetMsg(msg int) {
	s.gameState.SetMsg(msg)
}

type GameMain struct {
	gameState GameState
}

func NewGameMain(game *Game) *GameMain {
	g := &GameMain{
		gameState: GameState{game: game},
	}
	f := NewField(game.playerData)
	f.LoadFieldData(field_data)
	game.field = f
	game.player.Initialize()
	return g
}

func (g *GameMain) Update() {
	g.gameState.game.field.Move()
	g.gameState.game.player.Move()
}

func (g *GameMain) Draw() {
	if g.gameState.game.playerData.lunkerMode {
		g.gameState.game.Draw("bg", 0, 0, 0, 240, 320, 240)
	} else {
		g.gameState.game.Draw("bg", 0, 0, 0, 0, 320, 240)
	}

	p := g.gameState.game.player.view.GetPosition()
	g.gameState.game.field.Draw(g.gameState.game, Position{X: int(p.X), Y: int(p.Y)})
	g.gameState.game.player.Draw(g.gameState.game)
}

func (g *GameMain) GetMsg() int {
	return g.gameState.GetMsg()
}

func (g *GameMain) SetMsg(msg int) {
	g.gameState.SetMsg(msg)
}

type IGameState interface {
	Update()
	Draw()
	GetMsg() int
	SetMsg(msg int)
}

type Game struct {
	gameState  IGameState
	field      *Field
	player     *Player
	playerData *PlayerData
	key        *Input
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
	g.player = NewPlayer(g)
	return ebiten.Run(g.Loop, g_width, g_height, 1, "Inovation 5")
}

func (g *Game) Loop(screen *ebiten.Image) error {
	g.screen = screen

	if g.gameState == nil {
		g.gameState = &TitleMain{gameState: GameState{game: g}}
	} else {
		switch g.gameState.GetMsg() {
		case GAMESTATE_MSG_REQ_TITLE:
			g.gameState = &TitleMain{gameState: GameState{game: g}}
			break
		case GAMESTATE_MSG_REQ_OPENING:
			g.gameState = NewOpeningMain(g)
			break
		case GAMESTATE_MSG_REQ_GAME:
			g.gameState = NewGameMain(g)
			break
		case GAMESTATE_MSG_REQ_ENDING:
			g.gameState = NewEndingMain(g)
			break
		case GAMESTATE_MSG_REQ_SECRET1:
			g.gameState = NewSecretMain(g, 1)
			break
		case GAMESTATE_MSG_REQ_SECRET2:
			g.gameState = NewSecretMain(g, 2)
			break
		}
	}
	g.gameState.Update()
	g.gameState.Draw()
	g.key.Update()
	return nil
}

func (g *Game) FillRect(x, y, w, h int, clr color.Color) error {
	return g.screen.DrawRect(x, y, w, h, clr)
}

type imgPart struct {
	px, py, sx, sy, sw, sh int
}

func (i *imgPart) Len() int                       { return 1 }
func (i *imgPart) Dst(_ int) (int, int, int, int) { return i.px, i.py, i.px + i.sw, i.py + i.sh }
func (i *imgPart) Src(_ int) (int, int, int, int) { return i.sx, i.sy, i.sx + i.sw, i.sy + i.sh }

func (g *Game) Draw(key string, px, py, sx, sy, sw, sh int) error {
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = &imgPart{px, py, sx, sy, sw, sh}
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
	game.key = &Input{}
	game.img = map[string]*ebiten.Image{}
	for _, f := range []string{"ino", "msg", "bg"} {
		var err error
		game.img[f], _, err = ebitenutil.NewImageFromFile(filepath.Join(imgDir, f + ".png"), ebiten.FilterNearest)
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
