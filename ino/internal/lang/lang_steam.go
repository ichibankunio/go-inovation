//go:build !js && steam
// +build !js
// +build steam

package lang

import (
	"os"

	"golang.org/x/text/language"

	"github.com/hajimehoshi/go-inovation/ino/internal/steamworks"
)

const appID = 1710390

func init() {
	if steamworks.RestartAppIfNecessary(appID) {
		os.Exit(1)
	}
	if !steamworks.Init() {
		panic("steamworks.Init failed")
	}
}

func SystemLang() language.Tag {
	switch steamworks.SteamApps().GetCurrentGameLanguage() {
	case "english":
		return language.English
	case "japanese":
		return language.Japanese
	}
	return language.English
}
