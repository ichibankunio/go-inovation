name=innovation2007
mkdir -p bin

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags=-H=windowsgui -o bin/${name}.exe .
