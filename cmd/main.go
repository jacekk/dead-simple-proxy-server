package main

import (
	"log"
	"os"
	"sync"

	"github.com/urfave/cli"

	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
	"github.com/jacekk/dead-simple-proxy-server/pkg/routing"
	"github.com/jacekk/dead-simple-proxy-server/pkg/worker"
)

var projectDir string

func init() {
	projectDir = helpers.GetProjectDir()
	helpers.LoadEnvs(projectDir)
}

func main() {
	runApp()
}

func runApp() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s", "serve"},
			Usage:   "Runs server",
			Action: func(c *cli.Context) error {
				return routing.InitRouter(os.Getenv("SERVER_PORT"))
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"w", "worker"},
			Usage:   "Runs worker",
			Action: func(c *cli.Context) error {
				return worker.Init()
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "server",
			Usage: "To init the server",
		},
		cli.BoolFlag{
			Name:  "worker",
			Usage: "To init the worker",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var wg sync.WaitGroup

		if ctx.Bool("server") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Fatal(routing.InitRouter(os.Getenv("SERVER_PORT")))
			}()
		}
		if ctx.Bool("worker") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Fatal(worker.Init())
			}()
		}
		wg.Wait()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
