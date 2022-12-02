package main

import (
	"os"

	"github.com/hilaily/lib/logrustool"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "modtool",
		Usage: "some tools to make golang module friendly",
		Commands: []*cli.Command{
			tagCommand(),
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "log",
				DefaultText: "false",
				Usage:       "print debug log",
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("log") {
				logrustool.SetLevel(logrus.DebugLevel)
			}
			return nil
		},
		Version: "v0.0.1",
	}

	if err := app.Run(os.Args); err != nil {
		pRed(err.Error())
	}
}
