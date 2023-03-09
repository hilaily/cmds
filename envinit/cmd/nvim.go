package cmd

import (
	"path/filepath"

	"github.com/fatih/color"
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
	res, err := exec.Run("brew install neovim")
	if err != nil {
		return err
	}
	color.Green(string(res))
	return n.config(ctx)
}

func (n *nvimCMD) config(ctx *cli.Context) error {
	return util.CheckLink(filepath.Join(util.HomeDir+"/.config/nvim"), filepath.Join(util.DotfileDir+"/nvim"))
}
