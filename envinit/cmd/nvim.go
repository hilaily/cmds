package cmd

import (
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/cmds/envinit/util"
)

var (
	NvimCMD = (&nvimCMD{}).cmd()
)

type nvimCMD struct{}

func (n *nvimCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  "nvim",
		Usage: "mange tool nvim",
		Subcommands: []*cli.Command{
			{Name: "install", Action: n.install},
			{Name: "config", Action: n.config},
		},
	}
}

func (n *nvimCMD) install(ctx *cli.Context) error {
	exec.MustRun("brew install neovim")
	n.config(ctx)
	return nil
}

func (n *nvimCMD) config(ctx *cli.Context) error {
	err := util.CheckLink(filepath.Join(util.HomeDir+"/.config/nvim"), filepath.Join(util.DotfileDir+"/nvim"))
	exec.CheckErr(err)
	return nil
}
