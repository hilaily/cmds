package exec

import "testing"

func TestRunWithOutput(t *testing.T) {
	RunWithOutput("wget https://golang.google.cn/dl/go1.20.2.linux-amd64.tar.gz")
}
