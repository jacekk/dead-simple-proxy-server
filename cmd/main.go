package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
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
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Runs migrations up or down",
			Action: func(c *cli.Context) error {
				return errors.New("@todo migrate command")
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s", "serve"},
			Usage:   "Runs server",
			Action: func(c *cli.Context) error {
				port := os.Getenv("SERVER_PORT")
				log.Printf("Server port: %s \n", port)

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
