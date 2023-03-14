package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/util"
)

var (
	ZSHCMD = (&zshCMD{}).cmd()
)

type zshCMD struct{}

func (z *zshCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  "zsh",
		Usage: "install zsh",
		Subcommands: []*cli.Command{
			{Name: "install", Action: z.install},
			{Name: "config", Action: z.config},
		},
	}
}

func (z *zshCMD) install(ctx *cli.Context) error {
	util.BrewInstall("zsh")
	z.config(ctx)
	return nil
}

func (z *zshCMD) config(ctx *cli.Context) error {
	util.SetupDotfile()
	return nil
}
