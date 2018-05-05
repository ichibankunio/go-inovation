package ino

import (
	"flag"
	"fmt"
	_ "image/png"
	"os"
	"runtime/pprof"

	"github.com/gopherjs/gopherjs/js"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/go-inovation/ino/internal/audio"
	"github.com/hajimehoshi/go-inovation/ino/internal/draw"
	"github.com/hajimehoshi/go-inovation/ino/internal/input"
)

type Game struct {
	resourceLoadedCh chan error
	scene            Scene
	gameData         *GameData
	cpup             *os.File
}

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func (g *Game) Loop(screen *ebiten.Image) error {
	// exp
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) && js.Global != nil {
		doc := js.Global.Get("document")
		canvas := doc.Call("getElementsByTagName", "canvas").Index(0)
		context := canvas.Call("getContext", "webgl")
		context.Call("getExtension", "WEBGL_lose_context").Call("loseContext")
		fmt.Println("Context Lost!")
		return nil
	}

	if g.resourceLoadedCh != nil {
		select {
		case err := <-g.resourceLoadedCh:
			if err != nil {
				return err
			}
			g.resourceLoadedCh = nil
		default:
		}
	}
	if g.resourceLoadedCh != nil {
		ebitenutil.DebugPrint(screen, "Now Loading...")
		return nil
	}

	input.Current().Update()

	if input.Current().IsKeyJustPressed(ebiten.KeyF) {
		f := ebiten.IsFullscreen()
		ebiten.SetFullscreen(!f)
		ebiten.SetCursorVisible(f)
	}

	if input.Current().IsKeyJustPressed(ebiten.KeyP) && *cpuProfile != "" && g.cpup == nil {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		g.cpup = f
		pprof.StartCPUProfile(f)
		fmt.Println("Start CPU Profiling")
	}

	if input.Current().IsKeyJustPressed(ebiten.KeyQ) && g.cpup != nil {
		pprof.StopCPUProfile()
		g.cpup.Close()
		g.cpup = nil
		fmt.Println("Stop CPU Profiling")
	}

	if g.scene == nil {
		g.scene = &TitleScene{}
	} else {
		switch g.scene.Msg() {
		case GAMESTATE_MSG_REQ_TITLE:
			audio.PauseBGM()
			g.scene = &TitleScene{}
		case GAMESTATE_MSG_REQ_OPENING:
			if err := audio.PlayBGM(audio.BGM1); err != nil {
				return err
			}
			g.scene = &OpeningScene{}
		case GAMESTATE_MSG_REQ_GAME:
			g.scene = NewGameScene(g)
		case GAMESTATE_MSG_REQ_ENDING:
			if err := audio.PlayBGM(audio.BGM1); err != nil {
				return err
			}
			g.scene = &EndingScene{}
		case GAMESTATE_MSG_REQ_SECRET1:
			if err := audio.PlayBGM(audio.BGM1); err != nil {
				return err
			}
			g.scene = NewSecretScene(1)
		case GAMESTATE_MSG_REQ_SECRET2:
			if err := audio.PlayBGM(audio.BGM1); err != nil {
				return err
			}
			g.scene = NewSecretScene(2)
		}
	}
	g.scene.Update(g)
	if ebiten.IsRunningSlowly() {
		return nil
	}
	g.scene.Draw(screen, g)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %.2f", ebiten.CurrentFPS()))
	return nil
}

func NewGame() (*Game, error) {
	game := &Game{
		resourceLoadedCh: make(chan error),
	}
	go func() {
		if err := draw.LoadImages(); err != nil {
			game.resourceLoadedCh <- err
			return
		}
		if err := audio.Load(); err != nil {
			game.resourceLoadedCh <- err
			return
		}
		close(game.resourceLoadedCh)
	}()
	if err := audio.Finalize(); err != nil {
		return nil, err
	}
	return game, nil
}
