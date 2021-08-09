name=innovation2007
mkdir -p bin

# Windows
# TODO: Should '-ldflags=-H=windowsgui' be added?
GOOS=windows GOARCH=amd64 go build -tags=steam -o bin/${name}.exe .

# macOS
# TODO: Create .app
go build -tags=steam -o bin/${name} .
