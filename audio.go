package main

import (
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

func initAudio() chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)

		const sampleRate = 44100
		audioContext = audio.NewContext(sampleRate)
		const soundDir = "resource/sound"
		for _, n := range soundFilenames {
			f, err := ebitenutil.OpenFile(filepath.Join(soundDir, n))
			if err != nil {
				ch <- err
				return
			}
			var s audio.ReadSeekCloser
			switch {
			case strings.HasSuffix(n, ".ogg"):
				var err error
				s, err = vorbis.Decode(audioContext, f)
				if err != nil {
					ch <- err
					return
				}
			case strings.HasSuffix(n, ".wav"):
				var err error
				s, err = wav.Decode(audioContext, f)
				if err != nil {
					ch <- err
					return
				}
			default:
				panic("invalid file name")
			}
			p, err := audioContext.NewPlayer(s)
			if err != nil {
				ch <- err
				return
			}
			soundPlayers[n] = p
		}
		ch <- nil
	}()
	return ch
}

func finalizeAudio() error {
	for _, p := range soundPlayers {
		if err := p.Close(); err != nil {
			return err
		}
	}
	return nil
}

type BGM string

const (
	BGM0 BGM = "ino1.ogg"
	BGM1 BGM = "ino2.ogg"
)

func SetBGMVolume(volume float64) {
	for _, b := range []BGM{BGM0, BGM1} {
		p := soundPlayers[string(b)]
		if !p.IsPlaying() {
			continue
		}
		p.SetVolume(volume)
		return
	}
}

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
	if err := StopBGM(); err != nil {
		return err
	}
	p := soundPlayers[string(bgm)]
	p.SetVolume(1)
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
