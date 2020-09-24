#!/bin/sh

APP="hugy.app"
mkdir -p $APP/Contents/{MacOS,Resources}
mkdir -p $APP/Contents/MacOS/bin
go build -o $APP/Contents/MacOS/hugy
cat > $APP/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>hugy</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>io.hby.hugy</string>
</dict>
</plist>
EOF
cp icons/icon.icns $APP/Contents/Resources/icon.icns
cp bin/hugo $APP/Contents/MacOS/bin/hugo
find $APP