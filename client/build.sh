#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

export os_all="linux windows darwin freebsd"
export arch_all="386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64"

for os in $os_all; do
    for arch in $arch_all; do
        echo building for $os/$arch
        target=build/tdp-blackbox-$os-$arch
        if [ x"$os" = x"windows" ]; then
            target="${target}.exe"
        fi
        GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o $target main.go
    done
done

####################################################################

for app in `ls build`; do
    gzip build/$app
done
