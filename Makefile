build:
	go build -o wg-connect cmd/cmd.go

server:
	./wg-connect server

client:
	./wg-connect client
