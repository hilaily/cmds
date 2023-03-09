package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/cmd"
)

func main() {
	app := &cli.App{
		Name:  "envinit",
		Usage: "to init a new development environment",
		Commands: []*cli.Command{
			cmd.BrewCMD,
			cmd.NvimCMD,
		},
		Version: "v0.0.1",
	}

	if err := app.Run(os.Args); err != nil {
		color.Red(err.Error())
	}
}
