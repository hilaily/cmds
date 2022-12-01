package main

import (
	e "os/exec"
	"strings"
)

func exec(cmd string) ([]byte, error) {
	cmds := strings.Fields(cmd)
	c := e.Command(cmds[0], cmds[1:]...)
	return c.CombinedOutput()
}
