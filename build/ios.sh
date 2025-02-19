set -e

gomobile bind -target ios -o ./Frameworks/Identity.xcframework -ldflags "-s -w"
