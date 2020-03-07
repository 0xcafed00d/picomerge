#!/bin/bash

mkdir picomerge_builds -p

cp README.md picomerge_builds
env GOOS=linux GOARCH=amd64 go build -o picomerge_builds/linux_64bit/picomerge main.go
env GOOS=linux GOARCH=arm go build -o picomerge_builds/linux_arm/picomerge main.go
env GOOS=darwin GOARCH=amd64 go build -o picomerge_builds/mac_64bit/picomerge main.go
env GOOS=windows GOARCH=386 go build -o picomerge_builds/win/picomerge.exe main.go

rm picomerge_builds.zip
zip -r picomerge_builds.zip picomerge_builds 