package ino

import (
	"syscall/js"

	"golang.org/x/text/language"

	"github.com/hajimehoshi/go-inovation/ino/internal/text"
)

func systemLang() language.Tag {
	nav := js.Global().Get("navigator")
	if !nav.Truthy() {
		return language.Japanese
	}
	str := nav.Get("language").String()
	newLang, _ := language.Parse(str)
	base, _ := newLang.Base()
	newLang, _ = language.Compose(base)
	for _, l := range text.Languages() {
		if newLang == l {
			return newLang
		}
	}
	return language.Japanese
}
