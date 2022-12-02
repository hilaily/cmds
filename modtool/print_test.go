package main

import (
	"testing"

	"github.com/hilaily/lib/logrustool"
	"github.com/sirupsen/logrus"
)

func TestPrintTable(t *testing.T) {
	logrustool.SetLevel(logrus.DebugLevel)
	data := []string{"README.md", "git_test.go", "mod.go", "tag.go",
		"cache.go", "go.mod", "mod_parse.go", "version.go",
		"exec.go", "go.sum", "mod_test.go",
		"exec_test.go", "log.go", "parse_test.go",
		"git.go", "main.go", "print.go", "a.go", "1.txt",
	}
	pTable(data)
}
