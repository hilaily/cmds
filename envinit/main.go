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
			cmd.ZCMD,
		},
		Version: "v0.0.1",
	}
	defer func() {
		r := recover()
		if r != nil {
			color.Red("%v", r)
		}
	}()

	_ = app.Run(os.Args)
}
