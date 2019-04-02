build: deps build_darwin

server:
	dist/wg-darwin server

client:
	dist/wg-darwin client

build_all: deps build_darwin build_linux

build_linux:
	GOOS=linux GOARCH=amd64 go build -o dist/wg-linux cmd/cmd.go

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o dist/wg-darwin cmd/cmd.go

deps:
	dep ensure
	mkdir -p dist
