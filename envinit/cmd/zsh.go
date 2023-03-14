package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/exec"
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
	exec.RunWithOutput("sudo apt install zsh")
	exec.RunWithOutput("sudo chsh -s $(which zsh)")
	z.config(ctx)
	return nil
}

func (z *zshCMD) config(ctx *cli.Context) error {
	exec.MustRun("wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O /tmp/zsh_install.sh")
	exec.MustRun("sh /tmp/zsh_install.sh")
	util.SetupDotfile()
	return nil
}
