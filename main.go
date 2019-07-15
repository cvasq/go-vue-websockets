package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var flags []cli.Flag

func init() {

	flags = []cli.Flag{
		cli.StringFlag{
			Name:   "listening-port",
			Value:  "8080",
			Usage:  "Listening Port",
			EnvVar: "LISTENING_PORT",
		},
	}
}

func main() {

	app := cli.NewApp()
	app.Usage = "Vue.js Websocket Example"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Carlos Vasquez",
			Email: "carlos@cobranix.com",
		},
	}
	app.Flags = flags

	app.Action = StartListener

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
