#!/bin/bash

mkdir picomerge_builds -p

cp README.md picomerge_builds
env GOOS=windows GOARCH=amd64 go build -o picomerge_builds/picomerge.exe main.go
env GOOS=linux GOARCH=amd64 go build -o picomerge_builds/picomerge_linux_64 main.go
env GOOS=linux GOARCH=arm go build -o picomerge_builds/picomerge_linux_arm main.go
env GOOS=darwin GOARCH=amd64 go build -o picomerge_builds/picomerge_macos_64 main.go

zip picomerge_builds.zip picomerge_builds/* 