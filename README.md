# Go port of "Inovation 2007" by Omega

http://o-mega.sakura.ne.jp/product/ino.html

# How to install and run

```
:; go get github.com/hajimehoshi/go-invation
:; cd $GOPATH/src/github.com/hajimehoshi/go-invation
:; go run main.go
```

# How to build for Android

At this directory, run

```
:; gomobile bind -target android -javapkg com.hajimehoshi.goinovation.go -o ./mobile/android/inovation/inovation.aar ./mobile
```

and run the Android studio project in `./mobile/android`.

# How to build for iOS

At this directory, run

```
:; gomobile bind -target ios -o ./mobile/ios/Inovation.framework ./mobile
```

and run the Xcode project in `./mobile/ios`.
