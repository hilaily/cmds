package cmd

import (
	"path/filepath"
	"runtime"

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
			{Name: "rawinstall", Action: n.rawInstall},
			{Name: "config", Action: n.config},
		},
	}
}

func (n *nvimCMD) install(ctx *cli.Context) error {
	exec.MustRun("brew install neovim")
	n.config(ctx)
	return nil
}

func (n *nvimCMD) rawInstall(ctx *cli.Context) error {
	url := "https://github.com/neovim/neovim/releases/download/stable/nvim-linux64.tar.gz"
	untar := "tar xzvf nvim-linux64.tar.gz"
	mv := "mv nvim-linux/ /usr/local/"
	os := runtime.GOOS
	if os == "darwin" {
		url = "https://github.com/neovim/neovim/releases/download/stable/nvim-macos.tar.gz"
		untar = "tar xzvf nvim-macos.tar.gz"
		mv = "mv nvim-macos/ /usr/local/"
	}
	if os == "window" {
		panic("not support")
		//url = "https://github.com/neovim/neovim/releases/download/stable/nvim-win64.zip"
		//untar = "unzip nvim-win64.zip"
	}

	exec.MustSHRun("cd /tmp/ && wget " + url)
	exec.MustSHRun(untar)
	exec.MustSHRun(mv)
	n.config(ctx)
	return nil
}

func (n *nvimCMD) config(ctx *cli.Context) error {
	util.SetupDotfile()
	err := util.CheckLink(filepath.Join(util.HomeDir+"/.config/nvim"), filepath.Join(util.DotfileDir+"/nvim"))
	exec.CheckErr(err)
	return nil
}
