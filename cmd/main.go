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
			Action: func(ctx *cli.Context) error {
				return routing.InitRouter()
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"w", "work"},
			Usage:   "Runs worker",
			Action: func(ctx *cli.Context) error {
				return worker.InitWorker()
			},
		},
		{
			Name:  "all",
			Usage: "Runs worker and server",
			Action: func(ctx *cli.Context) error {
				var wg sync.WaitGroup

				go func() {
					defer wg.Done()
					log.Fatal(routing.InitRouter())
				}()
				go func() {
					defer wg.Done()
					log.Fatal(worker.InitWorker())
				}()

				wg.Add(2)
				wg.Wait()

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
