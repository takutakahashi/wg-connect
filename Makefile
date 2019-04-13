OS := $(shell uname -s |tr '[A-Z]' '[a-z]')

build: deps
	GOOS=$(OS) GOARCH=amd64 go build -o dist/$(OS) cmd/cmd.go

server:
	dist/$(OS) server

client:
	dist/$(OS) client

sdp:
	dist/$(OS) sdp

deps:
	dep ensure
	protoc --proto_path=pkg/proto/src --go_out=pkg/proto/dist pkg/proto/src/sdp_exchange.proto
	mkdir -p dist
