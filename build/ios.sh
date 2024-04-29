set -e

/usr/local/go/bin/go get -u golang.org/x/mobile/bind

/Users/user/go/bin/gomobile bind -target ios -o ./Frameworks/Identity.xcframework
