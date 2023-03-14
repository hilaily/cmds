package cmd

import (
	"fmt"

	"github.com/hilaily/kit/pathx"
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/cmds/envinit/util"
)

var (
	TmuxCMD = (&tmuxCMD{name: "tmux"}).cmd()
)

type tmuxCMD struct {
	name string
}

func (g *tmuxCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  g.name,
		Usage: "install " + g.name,
		Subcommands: []*cli.Command{
			{Name: "install", Action: g.install},
			{Name: "config", Action: g.config},
		},
	}
}

func (g *tmuxCMD) install(ctx *cli.Context) error {
	exec.RunWithOutput("sudo apt install tmux")
	g.config(ctx)
	if !pathx.IsExist(util.HomeDir + "/.tmux/plugins/tmux-resurrect") {
		cmd := fmt.Sprintf("git clone --depth=1 https://github.com/tmux-plugins/tmux-resurrect %s/.tmux/plugins/tmux-resurrect", util.HomeDir)
		exec.MustRun(cmd)
	}
	g.config(ctx)
	return nil
}

func (g *tmuxCMD) config(ctx *cli.Context) error {
	util.CheckLink(util.HomeDir+"/.tmux.conf", util.HomeDir+"/.dotfile/.tmux.conf")
	return nil
}
