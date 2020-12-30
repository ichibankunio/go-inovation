package assets

//go:generate file2byteslice -package assets -input=./resources/images/color/bg.png -output=image.bg.png.go -var=ImageBg
//go:generate file2byteslice -package assets -input=./resources/images/color/ino.png -output=image.ino.png.go -var=ImageIno
//go:generate file2byteslice -package assets -input=./resources/images/color/msg_en.png -output=image.msg_en.png.go -var=ImageMsgEn
//go:generate file2byteslice -package assets -input=./resources/images/color/msg_ja.png -output=image.msg_ja.png.go -var=ImageMsgJa
//go:generate file2byteslice -package assets -input=./resources/images/color/touch.png -output=image.touch.png.go -var=ImageTouch
//go:generate file2byteslice -package assets -input=./resources/sound/damage.wav -output=sound.damage.wav.go -var=SoundDamage
//go:generate file2byteslice -package assets -input=./resources/sound/heal.wav -output=sound.heal.wav.go -var=SoundHeal
//go:generate file2byteslice -package assets -input=./resources/sound/ino1.ogg -output=sound.ino1.ogg.go -var=SoundIno1
//go:generate file2byteslice -package assets -input=./resources/sound/ino2.ogg -output=sound.ino2.ogg.go -var=SoundIno2
//go:generate file2byteslice -package assets -input=./resources/sound/itemget.wav -output=sound.itemget.wav.go -var=SoundItemget
//go:generate file2byteslice -package assets -input=./resources/sound/itemget2.wav -output=sound.itemget2.wav.go -var=SoundItemget2
//go:generate file2byteslice -package assets -input=./resources/sound/jump.wav -output=sound.jump.wav.go -var=SoundJump
