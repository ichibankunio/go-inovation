package main

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/exp/audio"
	"github.com/hajimehoshi/ebiten/exp/audio/vorbis"
	"github.com/hajimehoshi/ebiten/exp/audio/wav"
)

var (
	audioContext   *audio.Context
	soundFilenames = []string{
		"damage.wav",
		"heal.wav",
		"ino1.ogg",
		"ino2.ogg",
		"itemget.wav",
		"itemget2.wav",
		"jump.wav",
	}
	soundPlayers = map[string]*audio.Player{}
)

func initAudio() error {
	const sampleRate = 44100
	audioContext = audio.NewContext(sampleRate)
	const soundDir = "resource/sound"
	for _, n := range soundFilenames {
		f, err := ebitenutil.OpenFile(filepath.Join(soundDir, n))
		// TODO: Should we close the file handler?
		if err != nil {
			return err
		}
		var s io.ReadSeeker
		switch {
		case strings.HasSuffix(n, ".ogg"):
			var err error
			s, err = vorbis.Decode(audioContext, f)
			if err != nil {
				return err
			}
		case strings.HasSuffix(n, ".wav"):
			var err error
			s, err = wav.Decode(audioContext, f)
			if err != nil {
				return err
			}
		default:
			panic("invalid file name")
		}
		p, err := audioContext.NewPlayer(s)
		if err != nil {
			return err
		}
		soundPlayers[n] = p
	}
	return nil
}

type BGM string

const (
	BGM0 BGM = "ino1.ogg"
	BGM1 BGM = "ino2.ogg"
)

func StopBGM() error {
	for _, b := range []BGM{BGM0, BGM1} {
		p := soundPlayers[string(b)]
		if err := p.Pause(); err != nil {
			return err
		}
		if err := p.Rewind(); err != nil {
			return err
		}
	}
	return nil
}

func PlayBGM(bgm BGM) error {
	StopBGM()
	p := soundPlayers[string(bgm)]
	return p.Play()
}

type SE string

const (
	SE_DAMAGE   SE = "damage.wav"
	SE_HEAL     SE = "heal.wav"
	SE_ITEMGET  SE = "itemget.wav"
	SE_ITEMGET2 SE = "itemget2.wav"
	SE_JUMP     SE = "jump.wav"
)

func PlaySE(se SE) error {
	p := soundPlayers[string(se)]
	if err := p.Rewind(); err != nil {
		return err
	}
	if err := p.Play(); err != nil {
		return err
	}
	return nil
}
