package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/buger/goterm"
	"github.com/sirupsen/logrus"
)

func pRed(format string, a ...any) {
	fmt.Printf(format, a...)
}

func pGreen(format string, a ...any) {
	fmt.Printf(format, a...)
}

func pNomarl(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func pTable(data []string) {
	width := goterm.Width()
	logrus.Debug("ssh width: ", width)
	maxData := len(data[0])
	for _, v := range data {
		if len(v) > maxData {
			maxData = len(v)
		}
	}
	table(data, int(width)/(maxData))
}

func table(data []string, cols int) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	l := len(data)
	for i, v := range data {
		fmt.Fprint(writer, v+"\t")
		if i == l-1 {
			break
		}
		if (i+1)%cols == 0 {
			fmt.Fprintln(writer)
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
