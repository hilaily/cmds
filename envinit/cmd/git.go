package cmd

import (
	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/cmds/envinit/util"
	"github.com/urfave/cli/v2"
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
	err := util.CheckLink(util.HomeDir+"/.gitconfig", util.HomeDir+"/.dotfile/git/.gitconfig")
	exec.CheckErr(err)
	return nil
}
