package util

import (
	"os"
)

var (
	HomeDir    = os.Getenv("HOME")
	DotfileDir = HomeDir + "/.dotfile"
)
