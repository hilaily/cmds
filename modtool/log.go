package main

import (
	"github.com/hilaily/lib/logrustool"
	"github.com/sirupsen/logrus"
)

func init() {
	logrustool.FormatOnlyMsg()
	logrustool.SetLevel(logrus.InfoLevel)
}
