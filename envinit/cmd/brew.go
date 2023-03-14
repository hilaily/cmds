package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/envinit/config"
	selfExec "github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/cmds/envinit/util"
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
			{Name: "config", Action: b.config},
		},
	}
}

func (b *brewCMD) install(ctx *cli.Context) error {
	file := "https://github.com/Homebrew/install/raw/HEAD/install.sh"
	if config.InCN {
		file = "https://mirrors.ustc.edu.cn/misc/brew-install.sh"
	}
	targetFile := "/tmp/brew_install.sh"
	selfExec.MustRun(fmt.Sprintf("wget %s -O %s", file, targetFile))

	env := os.Environ()
	//env := append(os.Environ(), `NONINTERACTIVE=1`)
	if config.InCN {
		env = append(env,
			`HOMEBREW_BREW_GIT_REMOTE=https://mirrors.ustc.edu.cn/brew.git`,
			`HOMEBREW_CORE_GIT_REMOTE=https://mirrors.ustc.edu.cn/homebrew-core.git`,
			`HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles`,
			`HOMEBREW_API_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles/api`,
		)
	}
	c := exec.Command("bash", targetFile)
	c.Env = env
	err := selfExec.RunCmdWithOutput(c)
	selfExec.CheckErr(err)
	b.config(ctx)
	return nil
}

func (b *brewCMD) config(ctx *cli.Context) error {
	util.SetupDotfile()
	return nil
}
