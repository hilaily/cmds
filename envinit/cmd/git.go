package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/util"
)

var (
	GitCMD = (&gitCMD{name: "git"}).cmd()
)

type gitCMD struct {
	name string
}

func (g *gitCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  g.name,
		Usage: "install " + g.name,
		Subcommands: []*cli.Command{
			{Name: "install", Action: g.install},
		},
	}
}

func (g *gitCMD) install(ctx *cli.Context) error {
	util.BrewInstall(g.name)
	g.config(ctx)
	return nil
}

func (g *gitCMD) config(ctx *cli.Context) error {
	util.SetupDotfile()
	return nil
}
