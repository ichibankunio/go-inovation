package audio

import (
	"strings"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/go-inovation/ino/internal/assets"
)

var (
	audioContext   *audio.Context
	soundFilenames = []string{
		"damage.wav",
		"heal.wav",
		"ino1.mp3",
		"ino2.mp3",
		"itemget.wav",
		"itemget2.wav",
		"jump.wav",
	}
	soundPlayers = map[string]*audio.Player{}
)

func init() {
	const sampleRate = 44100
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		panic(err)
	}
}

func Load() error {
	for _, n := range soundFilenames {
		b, err := assets.Asset("resources/sound/" + n)
		if err != nil {
			return err
		}
		f := audio.BytesReadSeekCloser(b)
		var s audio.ReadSeekCloser
		switch {
		case strings.HasSuffix(n, ".mp3"):
			stream, err := mp3.Decode(audioContext, f)
			if err != nil {
				return err
			}
			s = audio.NewInfiniteLoop(stream, stream.Length())
		case strings.HasSuffix(n, ".wav"):
			stream, err := wav.Decode(audioContext, f)
			if err != nil {
				return err
			}
			s = stream
		default:
			panic("invalid file name")
		}
		p, err := audio.NewPlayer(audioContext, s)
		if err != nil {
			return err
		}
		soundPlayers[n] = p
	}
	return nil
}

func Finalize() error {
	for _, p := range soundPlayers {
		if err := p.Close(); err != nil {
			return err
		}
	}
	return nil
}

type BGM string

const (
	BGM0 BGM = "ino1.mp3"
	BGM1 BGM = "ino2.mp3"
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

func PauseBGM() {
	for _, b := range []BGM{BGM0, BGM1} {
		p := soundPlayers[string(b)]
		p.Pause()
	}
}

func ResumeBGM(bgm BGM) {
	PauseBGM()
	p := soundPlayers[string(bgm)]
	p.SetVolume(1)
	p.Play()
}

func PlayBGM(bgm BGM) error {
	PauseBGM()
	p := soundPlayers[string(bgm)]
	p.SetVolume(1)
	if err := p.Rewind(); err != nil {
		return err
	}
	p.Play()
	return nil
}

type SE string

const (
	SE_DAMAGE   SE = "damage.wav"
	SE_HEAL     SE = "heal.wav"
	SE_ITEMGET  SE = "itemget.wav"
	SE_ITEMGET2 SE = "itemget2.wav"
	SE_JUMP     SE = "jump.wav"
)

func PlaySE(se SE) {
	p := soundPlayers[string(se)]
	p.Rewind()
	p.Play()
}
