// +build darwin,!arm,!arm64 freebsd linux windows
// +build !android
// +build !ios

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/go-inovation/ino"
)

var (
	memProfile = flag.String("memprofile", "", "write memory profile to file")
	traceOut   = flag.String("trace", "", "write trace to file")
)

func main() {
	flag.Parse()

	if *traceOut != "" {
		f, err := os.Create(*traceOut)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		trace.Start(f)
		defer trace.Stop()
	}

	game, err := ino.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.Run(game.Loop, ino.ScreenWidth, ino.ScreenHeight, ino.Scale(), ino.Title); err != nil {
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
