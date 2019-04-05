OS := $(shell uname -s |tr '[A-Z]' '[a-z]')

build: deps
	GOOS=$(OS) GOARCH=amd64 go build -o dist/wg-$(OS) cmd/cmd.go

server:
	dist/wg-$(OS) server

client:
	dist/wg-$(OS) client

deps:
	dep ensure
	mkdir -p dist
