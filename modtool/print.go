package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/buger/goterm"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func pRed(format string, a ...any) {
	color.Red(format, a...)
}

func pBlue(format string, a ...any) {
	color.Blue(format, a...)
}

func pNomarl(format string, a ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf(format, a...)
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
