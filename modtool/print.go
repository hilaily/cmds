package main

import "github.com/sirupsen/logrus"

func pRed(format string, a ...any) {
	logrus.Errorf(format, a...)
}

func pGreen(format string, a ...any) {
	logrus.Infof(format, a...)
}

func pNomarl(format string, a ...any) {
	logrus.Printf(format+"\n", a...)
}
