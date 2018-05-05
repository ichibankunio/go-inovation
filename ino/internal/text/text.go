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
	TextIDEnding
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
はたすべき　<red>しめい</red>なのだ！



わたしの　いのししとしての　<red>ちが　さわぐ</red>！`,
		TextIDEnding: `あつめた、<red>じんぎ</red>たちが　かがやきだす！



うおーっ！

わたしは　さけびごえを　あげ
ひかりの　なかへ…

ほっぷ　すてっぷ　じゃんぷ
かーるいす！



やがて　ひとつの　いんしは
その　いしを　もとの　ばしょへ
かいきさせ
きおくの しんえんに　きざまれた
きげんの　いしきを
おもい　おこさせる　だろう…





2007　しんねん　はじまる！
<red>あけましておめでとう！</red>





<red>――</red>　ここから　くれじっと　<red>――</red>


いのべーじょん2007
＊＊ くれじっとの なかに いる ＊＊



すーぱー　ざつよう　にんげん
おめが　／　（　゜ワ゜ノ

おんがく　にんげん
どん

ふいーるど　まっぷ　にんげん
おめが　／　（　゜ワ゜ノ
げっく
ずい
３５１

たいとる　めいめい　にんげん
わんきち

てすとぷれい　にんげん
げーむへる2000
げっく
３５１
ずぃ
あかじゃ

すべさる　さんくす　にんげん
げーむ　せいさく　ぎじゅつ　いた
どにちまでに　げーむを　つくる　すれ

HTML5　いしょく
はねだ

ごー　げんご　いしょく
ほしはじめ


ふろどうーすど　ばい
おめが



<red>――</red>　くれじっと　ここまで　<red>――</red>

えんどう　おふろに　はいる
`,
	},
}

func Get(lang language.Tag, id TextID) string {
	return texts[lang][id]
}
