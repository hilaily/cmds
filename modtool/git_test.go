package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitRepoPath(t *testing.T) {
	g := &git{}
	data := []struct {
		origin string
		exp    string
	}{
		{origin: "git@github.com:hilaily/cmds.git", exp: "github.com/hilaily/cmds"},
		{origin: "https://github.com/hilaily/cmds.git", exp: "github.com/hilaily/cmds"},
	}
	for _, v := range data {
		res := g.getRepoPath(v.origin)
		assert.Equal(t, v.exp, res)
	}
}
