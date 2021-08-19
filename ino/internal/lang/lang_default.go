//go:build !js && !steam
// +build !js,!steam

package lang

import (
	"golang.org/x/text/language"
)

func SystemLang() language.Tag {
	// TODO: Implement this correctly
	return language.English
}
