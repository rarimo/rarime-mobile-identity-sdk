set -e

# go get golang.org/x/mobile/cmd/gomobile@v0.0.0-20250210185054-b38b8813d607

gomobile bind -target ios -o ./Frameworks/Identity.xcframework -ldflags "-s -w"
