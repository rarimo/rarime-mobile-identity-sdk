set -e

go get -u golang.org/x/mobile/bind

export PATH=$PATH:~/go/bin
gomobile init

gomobile bind -target=android/arm64 -o Identity.aar -ldflags="-s -w" -trimpath