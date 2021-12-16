#!/usr/bin/env bash
clear
echo "Building cnsbackup-config"
echo "------------------"
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR"

echo "Cleaning up previous builds and packages"
rm -rf bin/*
rm -rf pkg/*


echo "Build executables per platform"
OUTPUT="bin/linux/amd64/cnsbackup-config"
echo " - linux-amd64 --> $OUTPUT"
GOOS=linux GOARCH=amd64 go build -o $OUTPUT main.go

OUTPUT="bin/windows/amd64/cnsbackup-config.exe"
echo " - windows-amd64 --> $OUTPUT"
GOOS=windows GOARCH=amd64 go build -o $OUTPUT main.go

OUTPUT="bin/darwin/amd64/cnsbackup-config"
echo " - darwin-amd64 --> $OUTPUT"
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT main.go

OUTPUT="bin/darwin/arm64/cnsbackup-config"
echo " - darwin-arm64 --> $OUTPUT"
GOOS=darwin GOARCH=arm64 go build -o $OUTPUT main.go

OUTPUT="bin/freebsd/amd64/cnsbackup-config"
echo " - freebsd-amd64 --> $OUTPUT"
GOOS=freebsd GOARCH=amd64 go build -o $OUTPUT main.go
