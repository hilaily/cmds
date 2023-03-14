package cmd

import (
	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/kit/pathx"
	"github.com/urfave/cli/v2"
)

var (
	ZCMD = (&zCMD{}).cmd()
)

type zCMD struct{}

func (z *zCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  "z",
		Usage: "install z",
		Subcommands: []*cli.Command{
			{Name: "install", Action: z.install},
		},
	}
}

func (z *zCMD) install(ctx *cli.Context) error {
	if !pathx.IsExist("/usr/local/z.lua") {
		exec.RunWithOutput("sudo git clone --depth=1 https://github.com/skywind3000/z.lua.git /usr/local/z.lua")
	}
	return nil
}
