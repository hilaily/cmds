package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hilaily/cmds/envinit/exec"
)

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
