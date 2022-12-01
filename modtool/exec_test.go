package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	res, err := exec("git remote get-url --all origin")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	t.Log(string(res))
}
