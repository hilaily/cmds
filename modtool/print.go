package main

import "fmt"

func pRed(format string, a ...any) {
	fmt.Printf(format, a...)
}

func pGreen(format string, a ...any) {
	fmt.Printf(format, a...)
}

func pNomarl(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
