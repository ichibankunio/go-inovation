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
	TextIDEndingScore1
	TextIDEndingScore2
	TextIDEndingScore3
	TextIDSecretCommand
	TextIDSecretClear
	TextIDItemPowerUp
	TextIDItemFuji
	TextIDItemBushi
	TextIDItemApple
	TextIDItemV
	TextIDItemTaka
	TextIDItemShoulder
	TextIDItemDagger
	TextIDItemKatakata
	TextIDItemNasu
	TextIDItemBonus
	TextIDItemNurse
	TextIDItemNazuna
	TextIDItemGameHell
	TextIDItemGundam
	TextIDItemPoed
	TextIDItemMilestone
	TextIDItem1Yen
	TextIDItemTriangle
	TextIDItemOmega
	TextIDItemLife
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

えんどう　おふろに　はいる`,
		TextIDEndingScore1: "せいせき　はぴょう",
		TextIDEndingScore2: "かくとく　あいてむ",
		TextIDEndingScore3: "くりあ　たいむ",
		TextIDSecretCommand: `たいとるで

ひだり ひだり ひだり
みぎ　みぎ　みぎ
ひだり　みぎ`,
		TextIDSecretClear: `おめでとう。

あなたは
ぎじゅつ　さる
です。`,
		TextIDItemPowerUp: `「みずぐすり」を　てに　いれた。
「てーれってれー」
「じゃんぶりょくが　あっぶ。
<red>くうちゅう　じゃんぶ</red>が　１かい
くわわる　くわわる！`,
		TextIDItemFuji: `「ふじ」を　てに　いれた。
じんぎの　ひとつ。
ふんかしない　かざん。
きゅうかざん　って
いうんだって。`,
		TextIDItemBushi: `「ぶし」を　てに　いれた。
じゃぱにーず　ないと。
あめりか　だいとうりょうも　ぶしどー。
さむらーい　さむらーい　ぶしどー。`,
		TextIDItemApple: `「ふじりんご」を　てに　いれた。
あかい　かじつ。
あーかーい　りんごーにー
くちびーる　よーせーてー。`,
		TextIDItemV: `「ぶい」を　てに　いれた。
たたかいの　さけび！
「しんかとは　ひとと　げーむの
がったいだ！」`,
		TextIDItemTaka: `「たか」を　てに　いれた。
じんぎの　ひとつ。
そらの　はんたー。
こだいあすてかでは
かみの　つかい　なんだ。`,
		TextIDItemShoulder: `「かた」を　てに　いれた。
どうたいの　うえ、
うでの　つけね。
かたが　あかいと
じかんに　おくれる。`,
		TextIDItemDagger: `「だがー」を　てに　いれた。
みじかい　けん
ひだりてたての　かわりに。
こがたなの　かわりに。
ぼうけんの　おともに　どうぞ。`,
		TextIDItemKatakata: `「かたかた」を　てに　いれた。
かたかた…
「これは　まるたーがいすと…」
「ちがう！ぷらずま　だ！！」`,
		TextIDItemNasu: `「なす」を　てに　いれた。
じんぎの　ひとつ。
むらさきに　かがやく
やさいの　おうさま
でも　あまり　たべたきが　しない。`,
		TextIDItemBonus: `「ぼうなす」を　てに　いれた。
あ　ぼうなす！
ふぇいたりてぃ　ぼうなす！
ぱしふぃすと　ぼうなす！
でぃす　いず　ざ　ぼうなす！`,
		TextIDItemNurse: `「なーす」を　てに いれた。
「かんごふでは　ない！
『かんごし』と　よべ！
この　ペいしえんと　どもめ！」`,
		TextIDItemNazuna: `「なずな」を　てに　いれた。
べつめい、ぺんぺんぐさ。
なずなが　とおったあとには
ぺんぺんぐさすら　のこらないという`,
		TextIDItemGameHell: `「げーむへる」を　てに　いれた。
ようこそ。
げーむ　せいさくしゃと
その　しゅうへんの　ための
こみゅにていへ。`,
		TextIDItemGundam: `「じっしゃがんだむ」を　てに　いれた。
「そうさは　かんたんだ。
せんとう　こんぴゅーたの
すいっちを　いれるだけで　いい。」`,
		TextIDItemPoed: `「ほえど」を　てに　いれた。
「ふん。
おもしろくなって
きやがったぜ！」`,
		TextIDItemMilestone: `「まいるまーく」を　てに　いれた。
さいしんさく「からす」
ぜっさん　かどうちゅう　でしゅー`,
		TextIDItem1Yen: `「いちえんさつ」を　てに　いれた。
「くらくて　よくみえないわ」
「ほうら　あかるいだろう」
「『くらくて　よくみえないわ』と
かいてある」`,
		TextIDItemTriangle: `「とらいあんぐる」を　てに　いれた。
すべって　ころんで　おおいたけん。
しゃちょうは　いま
どうしているのか…`,
		TextIDItemOmega: `「おめがの　くんしょう」を　てに　いれた。
こんな　げーむに
まじに　なって
どうも　ありがとう`,
		TextIDItemLife: `「はーとの　うつわ」を　てに　いれた。
でれででーん！
<red>らいふ</red>の　<red>じょうげん</red>を
１ふやしてあげる
ああ、なんて　たくましいの…`,
	},
}

func Get(lang language.Tag, id TextID) string {
	return texts[lang][id]
}
