//go:build !js
// +build !js

package ino

import (
	"golang.org/x/text/language"
)

func systemLang() language.Tag {
	return language.Japanese
}
