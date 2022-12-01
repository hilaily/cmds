package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	l := len(os.Args)
	if l < 3 {
		logrus.Error("arguments wrong")
	}
	typ := os.Args[1]
	cmd := os.Args[2]
	do(typ, cmd)
}

func do(typ string, cmd string) {
	switch typ {
	case "tag":
		t, err := newTag()
		if err != nil {
			pRed(err.Error())
			return
		}
		t.do(cmd, os.Args[3:]...)
	default:
		logrus.Errorf("type is wrong")
	}
}
