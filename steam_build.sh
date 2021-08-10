name=innovation2007
name_app=Innovation2007.app
mkdir -p bin

# Windows
# TODO: Should '-ldflags=-H=windowsgui' be added?
GOOS=windows GOARCH=amd64 go build -tags=steam -o bin/${name}.exe .

# macOS
rm -rf bin/${name_app}
mkdir -p bin/${name_app}/Contents/MacOS
mkdir -p bin/${name_app}/Contents/Resources

go build -tags=steam -o bin/${name_app}/Contents/MacOS/${name} .
cp steam/icon/icon_512x512.icns bin/${name_app}/Contents/Resources/icon.icns
echo '<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>CFBundleExecutable</key>
    <string>{{.Name}}</string>
    <key>CFBundleIconFile</key>
    <string>icon.icns</string>
    <key>NSHighResolutionCapable</key>
    <true />
  </dict>
</plist>' | sed -e "s/{{.Name}}/${name}/g" > bin/${name_app}/Contents/Info.plist
