package main

import (
	"github.com/takutakahashi/wg-connect/pkg/client"
	"github.com/takutakahashi/wg-connect/pkg/sdp"
	"github.com/takutakahashi/wg-connect/pkg/server"
	"github.com/urfave/cli"
	"os"
)

var Commands = []cli.Command{
	commandClient,
	commandServer,
	commandSDP,
}

var commandClient = cli.Command{
	Name:        "client",
	Usage:       "usage",
	Description: "desc",
	Action:      doClient,
}

var commandServer = cli.Command{
	Name:        "server",
	Usage:       "usage",
	Description: "desc",
	Action:      doServer,
}

var commandSDP = cli.Command{
	Name:        "sdp",
	Usage:       "usage",
	Description: "desc",
	Action:      doSDP,
}

func doClient(c *cli.Context) {
	client.ConnectServer()
}

func doServer(c *cli.Context) {
	server.Start()

}

func doSDP(c *cli.Context) {
	sdp.StartServer()

}

var Version string = "0.9.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "wgconnect"
	app.Usage = "wireguard connect assist for NAT"
	app.Version = Version
	app.Author = "takutakahashi"
	app.Email = "taku.takahashi120@gmail.com"
	app.Commands = Commands
	return app
}
