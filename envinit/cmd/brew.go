package cmd

import (
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/config"
	selfExec "github.com/hilaily/cmds/envinit/exec"
)

var (
	BrewCMD = (&brewCMD{}).cmd()
)

type brewCMD struct{}

func (b *brewCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  "brew",
		Usage: "mange tool brew",
		Subcommands: []*cli.Command{
			{Name: "install", Action: b.install},
		},
	}
}

func (b *brewCMD) install(ctx *cli.Context) error {
	cmd := []string{"/bin/bash", "-c", "$(curl -fsSL https://github.com/Homebrew/install/raw/HEAD/install.sh)"}
	env := os.Environ()
	if config.InCN {
		cmd = []string{"/bin/bash", "-c", "$(curl -fsSL https://mirrors.ustc.edu.cn/misc/brew-install.sh)"}
		env = append(env,
			`HOMEBREW_BREW_GIT_REMOTE="https://mirrors.ustc.edu.cn/brew.git"`,
			`HOMEBREW_CORE_GIT_REMOTE="https://mirrors.ustc.edu.cn/homebrew-core.git"`,
			`HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"`,
			`HOMEBREW_API_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles/api"`,
		)
	}
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Env = env
	err := selfExec.RunCmdWithOutput(c)
	selfExec.CheckErr(err)
	return nil
}
