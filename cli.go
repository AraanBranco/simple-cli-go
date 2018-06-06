package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "simpleCLI"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "provider, p",
			Usage: "Call the provider 'aws'",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "aws",
			Usage: "commands for AWS",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
