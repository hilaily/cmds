package main

import (
	"github.com/hilaily/lib/cmdx"
	"github.com/hilaily/lib/logrustool"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/modtool/service"
)

func main() {
	app := &cli.App{
		Name:  "modtool",
		Usage: "some tools to make golang module friendly",
		Commands: []*cli.Command{
			tagCommand(),
			service.RenameCommand(),
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
		Version: "v0.2.0",
	}

	cmdx.WrapCli(app)
}
