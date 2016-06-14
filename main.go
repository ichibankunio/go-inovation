// +build !android

package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/go-inovation/ino"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	game, err := ino.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.Run(game.Loop, ino.ScreenWidth, ino.ScreenHeight, 2, ino.Title); err != nil {
		panic(err)
	}
}
