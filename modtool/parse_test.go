package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMod(t *testing.T) {
	res, err := findModFile()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestParse(t *testing.T) {
	res, err := getModName("./go.mod")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	t.Log(res)
}
