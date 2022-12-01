package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMod(t *testing.T) {
	m := &mod{}
	res, err := m.findModFile()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestParse(t *testing.T) {
	m := &mod{}
	res, err := m.getModName("./go.mod")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	t.Log(res)
}
