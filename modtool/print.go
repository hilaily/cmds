package main

import "log"

func pRed(format string, a ...any) {
	log.Printf(format, a...)
}

func pGreen(format string, a ...any) {
	log.Printf(format, a...)
}

func pNomarl(format string, a ...any) {
	log.Printf(format+"\n", a...)
}
