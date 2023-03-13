package exec

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Run(cmd string) (string, error) {
	cmds := strings.Fields(cmd)
	c := exec.Command(cmds[0], cmds[1:]...)
	res, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s %w", string(res), err)
	}
	return string(res), nil
}

func MustRun(cmd string) {
	res, err := Run(cmd)
	CheckErr(err)
	color.Green(string(res))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
