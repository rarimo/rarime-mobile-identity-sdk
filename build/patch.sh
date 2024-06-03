set -e

version=$(go version | awk '{print $3}' | tr -d "go")

if [ "$version" != "1.22.3" ]; then
    echo "Go version is not 1.22.3 or go is not installed"
    exit 1
fi

if [ ! -d "go" ]; then
    git clone git@github.com:golang/go.git
    git reset --hard adbfb672ba485630d75f8b5598228a63f4af08a4
fi

cd go/src

if [ ! -f "../pkg/tool/darwin_arm64/cgo" ]; then
    ./make.bash
fi

mv ../pkg/tool/darwin_arm64/cgo /usr/local/go/pkg/tool/darwin_arm64/cgo
