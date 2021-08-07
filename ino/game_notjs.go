//go:build !js
// +build !js

package ino

import (
	"golang.org/x/text/language"
)

func tryLoseContext() bool {
	return false
}

func systemLang() language.Tag {
	return language.Japanese
}
