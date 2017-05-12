// +build darwin,!arm,!arm64 freebsd linux windows
// +build !android
// +build !ios

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/go-inovation/ino"
)

var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile = flag.String("memprofile", "", "write memory profile to file")
)

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
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
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(fmt.Sprintf("could not write memory profile: %s", err))
		}
	}
}
