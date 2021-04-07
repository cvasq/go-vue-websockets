package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var flags []cli.Flag

func init() {

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "listening-port",
			Value:   "8080",
			Usage:   "Server Listening Port",
			EnvVars: []string{"LISTENING_PORT"},
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Usage = "Vue.js Websocket Client Example"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Flags = flags

	app.Action = StartListener

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
