# Go port of "Inovation 2007" by Omega

http://o-mega.sakura.ne.jp/product/ino.html

# How to install and run

```
:; go get github.com/hajimehoshi/go-invation
:; cd $GOPATH/src/github.com/hajimehoshi/go-invation
:; go run *.go
```

# How to build for Android

```
:; gomobile bind -javapkg com.hajimehoshi.goinovation.go -o /path/to/inovation.aar .
```

# How to build for iOS

```
:; gomobile bind -target ios -o /path/to/Inovation.framework .
```
