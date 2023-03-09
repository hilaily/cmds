package util

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hilaily/cmds/envinit/exec"
)

func CheckLink(softLink, src string) error {
	f, err := os.Stat(softLink)
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
	exec.Run(fmt.Sprintf("ln -s %s %s", src, softLink))
	return nil
}
