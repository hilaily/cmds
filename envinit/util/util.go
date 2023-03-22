package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hilaily/cmds/envinit/exec"
	"github.com/hilaily/kit/pathx"
)

func SetupDotfile() {
	tgt := filepath.Join(HomeDir, "/.dotfile")
	if !pathx.IsExist(tgt) {
		exec.MustSHRun("git clone https://github.com/hilaily/.dotfile.git "+tgt, "GIT_SSL_NO_VERIFY=true")

		exec.MustSHRun("rm -f ~/.zshrc")
		exec.MustSHRun(`ln -sf ~/.dotfile/.zshrc ~/`)
		exec.MustSHRun(`mkdir -p ~/.oh-my-zsh/themes/`)
		exec.MustSHRun(`ln -sf ~/.dotfile/laily.zsh-theme ~/.oh-my-zsh/themes/`)
		exec.MustSHRun(`ln -sf ~/.dotfile/.tmux.conf ~/`)
		exec.MustSHRun(`ln -sf ~/.dotfile/git/.gitconfig ~/.gitconfig`)
	}
}

func BrewInstall(name string) {
	if !exec.CommandIsExist(name) {
		exec.MustRun("brew install " + name)
	}
}

func CheckLink(softLink, src string) error {
	f, err := os.Lstat(softLink)
	exist := true
	if err != nil {
		exist = os.IsExist(err)
	}
	if exist {
		if f.Mode()&os.ModeSymlink == os.ModeSymlink {
			real, err := os.Readlink(softLink)
			if err != nil {
				return err
			}
			if real == src {
				color.Green("path is already exist: %s", softLink)
			} else {
				color.Green("path is link to: %s", real)
			}
			return nil
		}
		color.Green("path is exist, move it as a backup, %s.bak", softLink)
		res, err := exec.Run(fmt.Sprintf("mv %s %s.bak", softLink, softLink))
		if err != nil {
			return err
		}
		color.Green(string(res))
	}
	color.Green("create soft link: %s", softLink)
	_, err = exec.Run("mkdir -p " + filepath.Dir(softLink))
	if err != nil {
		return err
	}
	_, err = exec.Run(fmt.Sprintf("ln -s %s %s", src, softLink))
	if err != nil {
		return err
	}
	return nil
}
