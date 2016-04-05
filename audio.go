package main

import (
	"path/filepath"
	"strings"
	"sync"

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

func (g *Game) loadAudio() {
	defer close(g.audioLoadedCh)


	type audioInfo struct {
		player *audio.Player
		key    string
	}
	ch := make(chan audioInfo)
	go func() {
		for p := range ch {
			soundPlayers[p.key] = p.player
		}
	}()

	const sampleRate = 44100
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		g.audioLoadedCh <- err
		return
	}
	var wg sync.WaitGroup
	for _, n := range soundFilenames {
		n := n
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err != nil {
				return
			}
			var f ebitenutil.ReadSeekCloser
			f, err = ebitenutil.OpenFile(filepath.Join("resource", "sound", n))
			if err != nil {
				return
			}
			var s audio.ReadSeekCloser
			switch {
			case strings.HasSuffix(n, ".ogg"):
				var stream *vorbis.Stream
				stream, err = vorbis.Decode(audioContext, f)
				if err != nil {
					return
				}
				s = audio.NewLoop(stream, stream.Size())
			case strings.HasSuffix(n, ".wav"):
				s, err = wav.Decode(audioContext, f)
				if err != nil {
					return
				}
			default:
				panic("invalid file name")
			}
			var p *audio.Player
			p, err = audioContext.NewPlayer(s)
			if err != nil {
				return
			}
			ch <- audioInfo{p, n}
		}()
	}
	wg.Wait()
	close(ch)
	g.audioLoadedCh <- err
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

func PauseBGM() error {
	for _, b := range []BGM{BGM0, BGM1} {
		p := soundPlayers[string(b)]
		if err := p.Pause(); err != nil {
			return err
		}
	}
	return nil
}

func ResumeBGM(bgm BGM) error {
	if err := PauseBGM(); err != nil {
		return err
	}
	p := soundPlayers[string(bgm)]
	p.SetVolume(1)
	return p.Play()
}

func PlayBGM(bgm BGM) error {
	if err := PauseBGM(); err != nil {
		return err
	}
	p := soundPlayers[string(bgm)]
	p.SetVolume(1)
	if err := p.Rewind(); err != nil {
		return err
	}
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
