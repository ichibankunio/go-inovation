package text

import (
	"golang.org/x/text/language"
)

type TextID int

const (
	TextIDStart TextID = iota
	TextIDStartLunker
	TextIDStartTouch
	TextIDOpening
)

var texts = map[language.Tag]map[TextID]string{
	language.Japanese: {
		TextIDStart:       "すぺーす　たたく　はじまる！",
		TextIDStartLunker: "らんかー　もーど　はじまる！",
		TextIDStartTouch:  "がめん　たっち　はじまる！",
		TextIDOpening: `めが　さめたら
<red>いのしし</red>に　なっていた。



おちつけ　おれ。
＊＊　ゆめの　なかにいる　＊＊



ゆめの　ちとゅう
どるいどの　よげんしゃが
わたしに　つげるのだった。



「<red>さんしゅの　じんぎ</red>を　みつけだせるのは
この　ゆめの　しゅじんこう…
そう　そなただけなのじゃ！」



いかなる　しょうがいをも　のりこえ
はつゆめ　おやくそくの
<red>さんしゅの　じんぎ</red>を　さがしだすこと…
それが　わたしの
はたすべき　<red>しめい</red>なのだ!



わたしの　いのししとしての　<red>ちが　さわぐ</red>！`,
	},
}

func Get(lang language.Tag, id TextID) string {
	return texts[lang][id]
}
