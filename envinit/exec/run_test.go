package exec

import "testing"

func TestRun(t *testing.T) {
	MustRun("ls $HOME")
}

func TestRunWithOutput(t *testing.T) {
	RunWithOutput("wget https://golang.google.cn/dl/go1.20.2.linux-amd64.tar.gz")
}
