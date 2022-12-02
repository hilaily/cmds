package main

import (
	"fmt"
	e "os/exec"
	"strings"
)

func exec(cmd string) ([]byte, error) {
	cmds := strings.Fields(cmd)
	c := e.Command(cmds[0], cmds[1:]...)
	res, err := c.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %w", string(res), err)
	}
	return res, nil
}
