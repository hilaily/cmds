package cmd

import (
	"path/filepath"

	"github.com/hilaily/kit/pathx"
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/cmds/envinit/util"
)

var (
	DotfileCMD = (&dotfileCMD{}).cmd()
)

type dotfileCMD struct{}

func (d *dotfileCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  "dotfile",
		Usage: "mange dotfile",
		Subcommands: []*cli.Command{
			{Name: "clone", Action: d.clone},
			{Name: "update", Action: d.update},
		},
	}
}

func (d *dotfileCMD) clone(ctx *cli.Context) error {
	util.SetupDotfile()
	return nil
}

func (d *dotfileCMD) update(ctx *cli.Context) error {
	if !pathx.IsExist(filepath.Join(util.HomeDir, "/.dotfile")) {
		util.SetupDotfile()
		return nil
	}
	exec.MustSHRun("cd ~/.dotfile")
	exec.MustRun("git pull origin master", "GIT_SSL_NO_VERIFY=true")
	return nil
}
