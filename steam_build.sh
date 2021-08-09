name=innovation2007
mkdir -p bin

# Windows
# GOOS=windows GOARCH=amd64 go build -ldflags=-H=windowsgui -tags=steam -o bin/${name}.exe .
GOOS=windows GOARCH=amd64 go build -tags=steam -o bin/${name}.exe .
