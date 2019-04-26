#! /bin/bash

if [ $# == 0 ]; then
  go build $(pwd)/main.go
elif [[ $1 == "--release" || "-r" ]]; then
  go build -ldflags="-s -w" $(pwd)/main.go
#elif [[ $1 == "--cross-compile" || "-cc" ]]; then
#  go build -o $(pwd)/main-linux-x64 $(pwd)/main.go
#  GOOS=darwin GOARCH=amd64 go build -o $(pwd)/main-darwin $(pwd)/main.go
#  GOOS=windows GOARCH=amd64 go build -o $(pwd)/main-windows-x64 $(pwd)/main.go
#elif [[ $1 == "--cross-compile" || "-cc" && $2 == "--release" || "-r" ]]; then
#  go build -o $(pwd)/main-linux-x64 -ldflags="-s -w" $(pwd)/main.go
#  GOOS=darwin GOARCH=amd64 go build -o $(pwd)/main-darwin -ldflags="-s -w" $(pwd)/main.go
#  GOOS=windows GOARCH=amd64 go build -o $(pwd)/main-windows-x64 -ldflags="-s -w" $(pwd)/main.go
#elif [[ $1 == "-ccr" ]]; then
#  go build -o $(pwd)/main-linux-x64 -ldflags="-s -w" $(pwd)/main.go
#  GOOS=darwin GOARCH=amd64 go build -o $(pwd)/main-darwin -ldflags="-s -w" $(pwd)/main.go
#  GOOS=windows GOARCH=amd64 go build -o $(pwd)/main-windows-x64 -ldflags="-s -w" $(pwd)/main.go
else
  echo "Not a valid argument!"
fi
